package main

import (
	"log"
	"os"
	"os/signal"
	"fmt"
	"flag"
	"net/http"
	"runtime"
	"time"

	ascf "github.com/zyfworks/AnotherSteamCommunityFix"
	"github.com/bitly/go-simplejson"
)

var (
	version    = "1.2.3"
	domainName = "steamcommunity.com"
	defaultIP = "184.26.221.151" // 日本東京都 Akamai CDN
	dnsList   = map[string]string{
		"OpenDNS_1":    "208.67.222.222:5353",
		"OpenDNS_2":    "208.67.220.220:443",
		"OpenDNS_2-fs": "208.67.220.123:443",
	}

	fixedIP string
)

func init() {
	flag.StringVar(&fixedIP, "ip", "", "手动指定IP地址")
	flag.Parse()
	log.SetFlags(log.Lshortfile)
}

func main() {
	sayHello()
	checkVersion()
	addHosts()

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Kill)
	signal.Notify(interrupt, os.Interrupt)

	log.Println("正在获取IP地址，请稍候~")
	var ipAddr string
	if len(fixedIP) == 0 {
		// name resolution by DNS
		address, err := ascf.DnsLookUp(domainName, dnsList)
		if err == nil && address != nil {
			ipAddr = address.String()
			log.Println("域名解析成功：", ipAddr)
			goto Start
		}
		// online lookup via http
		address, err = ascf.HttpLookup(domainName)
		if err == nil && address != nil {
			ipAddr = address.String()
			log.Println("获取IP地址成功：", ipAddr)
			goto Start
		}
		// using preset ip address
		ipAddr = defaultIP
		log.Println("获取IP地址失败，使用备用IP地址：", ipAddr)
	} else {
		ipAddr = fixedIP
		log.Println("使用手动指定的IP地址：", ipAddr)
	}

Start:
	fmt.Println("\n程序已启动，请不要关闭该窗口！")
	fmt.Println()
	fmt.Println("对于Mac和Linux用户，使用nohup命令运行程序可使其在后台运行。\n" +
		"\t└─ 在终端中进入程序所在目录后执行 “nohup sudo ./ascf &”即可。")

	go ascf.StartServingHTTPRedirect(http.StatusFound)
	go ascf.StartServingTCPProxy(":443", ipAddr+":443")

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
