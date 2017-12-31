package AnotherSteamCommunityFix

import (
	"errors"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/miekg/dns"
	"github.com/bitly/go-simplejson"
)

const RETRY = 10

func DnsLookUp(domainName, dnsServer string) (net.IP, error) {
	client := &dns.Client{}
	msg := &dns.Msg{}
	msg.SetQuestion(domainName+".", dns.TypeA)

	for i := 1; i <= RETRY; i++ {
		r, t, err := client.Exchange(msg, dnsServer)
		if err == nil && len(r.Answer) > 0 {
			log.Printf("域名解析耗时 %v\n", t)
			return r.Answer[0].(*dns.A).A, nil
		}
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
	for i, _ := range arr {
		ans := json.GetIndex(i)
		ip, err = ans.Get("value").String()
		if err == nil {
			return net.ParseIP(ip), nil
		}
	}

	return nil, errors.New("could not get the IP address")
}
