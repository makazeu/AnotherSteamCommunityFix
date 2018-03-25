package AnotherSteamCommunityFix

import (
	"log"
	"os"
)

var fatalLogger *log.Logger

func init() {
	fatalLogger = log.New(os.Stderr, "\n程序发生错误，已退出！\n", log.Lshortfile)
}
