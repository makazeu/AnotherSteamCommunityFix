package main

import (
	ascf "github.com/zyfworks/AnotherSteamCommunityFix"
	"log"
	"os"
	"os/signal"
	"fmt"
	"flag"
	"net/http"
)

const (
	Lredirect = iota + 1
	Lproxy
)

var (
	version    = "1.0"
	domainName = "steamcommunity.com"
	dnsServer  = "208.67.222.222:5353"
	defaultIP  = "104.115.125.124"

	mode                 int
	fixedIP              string
	chainNode, serveNode ascf.StringList
)

func init() {
	flag.IntVar(&mode, "mode", Lredirect, "1-转发模式、2-代理模式")
	flag.StringVar(&fixedIP, "ip", defaultIP, "备用的IP地址")
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	sayHello()
	addHosts()

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)

	if mode == Lredirect {
		var ipAddr string
		fmt.Println("程序设定为转发模式")

		address, err := ascf.LookUp(domainName, dnsServer, 10)
		if err != nil {
			ipAddr = fixedIP
			log.Println("域名解析失败，使用备用IP地址：", ipAddr)
		} else {
			ipAddr = address[0].String()
			log.Println("域名解析成功：", ipAddr)
		}

		go ascf.StartServingHTTPRedirect(http.StatusFound)
		go ascf.StartServingTCPProxy(":443", ipAddr+":443")
	} else if mode == Lproxy {
		fmt.Println("程序设定为代理模式")

		serveNode = append(serveNode, "tcp://:80/"+domainName+":80")
		serveNode = append(serveNode, "tcp://:443/"+domainName+":443")
		chainNode = append(chainNode, ProxyServer)
		var routes = ascf.NewGost(chainNode, serveNode)
		go routes.StartGostServing()
	} else {
		log.Fatal("illegal argument")
	}

	fmt.Println("程序已经启动，正在监听80和443端口，现在可正常访问Steam社区！")
	select {
	case <-interrupt:
		removeHosts()
	}
}

func addHosts() {
	if err := ascf.AddHosts("127.0.0.1", domainName); err != nil {
		log.Fatal(err)
	}
}

func removeHosts() {
	if err := ascf.RemoveHosts("127.0.0.1", domainName); err != nil {
		log.Fatal(err)
	}
}

func sayHello() {
	fmt.Printf("~ 欢迎使用AnotherSteamCommunityFix v%s ~\n", version)
	fmt.Println("Author: Makazeu [ Steam: Makazeu | Weibo: @Makazeu ]")
	fmt.Println()
}
