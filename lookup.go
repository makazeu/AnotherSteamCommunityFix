package AnotherSteamCommunityFix

import (
	"errors"
	"github.com/miekg/dns"
	"log"
	"net"
)

func LookUp(domainName, dnsServer string, retry int) ([]net.IP, error) {
	client := &dns.Client{}
	msg := &dns.Msg{}
	msg.SetQuestion(domainName+".", dns.TypeA)

	var address []net.IP
	for i := 1; i <= retry; i++ {
		r, t, err := client.Exchange(msg, dnsServer)
		if err == nil && len(r.Answer) != 0 {
			log.Printf("域名解析耗时 %v\n", t)
			for _,ans := range r.Answer  {
				ip := ans.(*dns.A).A
				address = append(address, ip)
			}
			return address, nil
		}
	}
	return nil, errors.New("域名解析失败")
}
