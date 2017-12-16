package main

import (
	ascf "github.com/zyfworks/AnotherSteamCommunityFix"
	"log"
	"os"
	"os/signal"
)

var (
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