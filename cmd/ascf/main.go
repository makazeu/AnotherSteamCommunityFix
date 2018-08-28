package main

import (
	"flag"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/zyfworks/libgost"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	ascf "github.com/zyfworks/AnotherSteamCommunityFix"
)

var (
	version    = "2.0.0"
	domainName = "steamcommunity.com"
)

func init() {
	flag.Parse()
	log.SetFlags(log.Lshortfile)

	libgost.InitGost()
}

func main() {
	sayHello()
	//checkVersion()
	addHosts()

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Kill)
	signal.Notify(interrupt, os.Interrupt)
	fmt.Println("程序已启动，请不要关闭该窗口！")
	fmt.Println()

	var chainNode, serveNode libgost.StringList
	serveNode = append(serveNode, "tcp://:443/"+domainName+":443")
	chainNode = append(chainNode, "kcp://51.15.131.212:20929")

	gost := libgost.NewGost(chainNode, serveNode)
	go gost.StartServing()

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

func checkVersion() {
	platform := runtime.GOOS + "_" + runtime.GOARCH
	client := &http.Client{Timeout: 5 * time.Second}
	r, err := client.Get("https://up.w21.win/update?" +
		"name=ascf&platform=" + platform + "&version=" + version)
	if err != nil {
		return
	}
	defer r.Body.Close()

	json, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		return
	}
	isOK, err := json.Get("ok").Bool()
	if err != nil || !isOK {
		return
	}

	newVersion, err := json.Get("version").String()
	if err != nil {
		return
	}
	if newVersion != version {
		fmt.Println("检测到新版本：" + newVersion + "，下载地址： https://steamcn.com/t339641-1-1")
	}
}
