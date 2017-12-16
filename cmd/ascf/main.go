package main

import (
	ascf "github.com/zyfworks/AnotherSteamCommunityFix"
	"log"
	"os"
	"os/signal"
	"fmt"
)

var (
	version = "0.1"
	domainName = "steamcommunity.com"
	chainNode, serveNode ascf.StringList
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	serveNode = append(serveNode, "tcp://:80/" + domainName + ":80")
	serveNode = append(serveNode, "tcp://:443/" + domainName + ":443")
	chainNode = append(chainNode, ProxyServer)
}

func main() {
	sayHello()
	addHosts()

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)

	var routes = ascf.NewGost(chainNode, serveNode)
	go routes.StartServing()

	select {
	case <- interrupt:
		removeHosts()
	}
}

func addHosts() {
	if err := ascf.AddHosts("127.0.0.1", domainName); err != nil {
		log.Fatal(err)
	}
}

func removeHosts()  {
	if err := ascf.RemoveHosts("127.0.0.1", domainName); err != nil {
		log.Fatal(err)
	}
}

func sayHello() {
	fmt.Printf("~ 欢迎使用AnotherSteamCommunityFix v%s ~\n", version)
	fmt.Println("Author: Makazeu [ Steam: Makazeu | Weibo: @Makazeu ]")
	fmt.Println()
	fmt.Println("程序已经启动，正在监听80和443端口，现在可正常访问Steam社区！")
}