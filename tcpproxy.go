package AnotherSteamCommunityFix

import (
	"net"
	"log"
	"time"
	"io"
)

func handleConn(conn net.Conn, remote string) {
	defer conn.Close()

	remoteConn, err := net.DialTimeout("tcp", remote, 15*time.Second)
	if err != nil {
		log.Println("dial remote error:", err)
		return
	}
	defer remoteConn.Close()

	go io.Copy(conn, remoteConn)
	io.Copy(remoteConn, conn)
}

func StartServingTCPProxy(local, remote string) {
	listener, err := net.Listen("tcp4", local)
	if err != nil {
		fatalLogger.Fatal("listen tcp error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fatalLogger.Fatal("accept tcp error:", err)
		}
		go handleConn(conn, remote)
	}
}
