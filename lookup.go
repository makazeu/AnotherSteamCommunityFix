package AnotherSteamCommunityFix

import (
	"errors"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/miekg/dns"
	"github.com/bitly/go-simplejson"
)

const RETRY = 5

func DnsLookUp(domainName string, dnsList map[string]string) (net.IP, error) {
	client := &dns.Client{Net: "tcp"}
	msg := &dns.Msg{}
	msg.SetQuestion(domainName+".", dns.TypeA)

	for dnsName, dnsAddress := range dnsList {
		for i := 1; i <= RETRY; i++ {
			r, t, err := client.Exchange(msg, dnsAddress)
			if err == nil && len(r.Answer) > 0 {
				log.Printf("使用%s解析域名成功，耗时: %v\n", dnsName, t)
				return r.Answer[rand.Int()%len(r.Answer)].(*dns.A).A, nil
			}
		}
		log.Printf("使用%s解析域名失败...", dnsName)
	}
	return nil, errors.New("域名解析失败")
}

func HttpLookup(domain string) (net.IP, error) {
	client := &http.Client{Timeout: 15 * time.Second}
	r, err := client.Get("https://dns-api.org/A/" + domain)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var ip string
	json, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		return nil, err
	}
	arr, err := json.Array()
	if err != nil {
		return nil, err
	}
	for i := range arr {
		ans := json.GetIndex(i)
		ip, err = ans.Get("value").String()
		if err == nil {
			return net.ParseIP(ip), nil
		}
	}

	return nil, errors.New("could not get the IP address")
}
