package AnotherSteamCommunityFix

import (
	"net"
	"log"
	"time"
)

func pipeAndClose(src, dst net.Conn)  {
	buf := make([]byte, 81920)
	for {
		n, err := src.Read(buf)
		if err != nil {
			break
		}

		if _, err := dst.Write(buf[:n]); err != nil {
			break
		}
	}
}

func handleConn(conn net.Conn, remote string)  {
	defer conn.Close()

	remoteConn, err := net.DialTimeout("tcp", remote, 15 * time.Second)
	if err != nil {
		log.Println("dial remote error:", err)
		return
	}
	defer remoteConn.Close()

	go pipeAndClose(remoteConn, conn)
	pipeAndClose(conn, remoteConn)
}

func StartServingTCPProxy(local, remote string)  {
	listener, err := net.Listen("tcp", local)
	if err != nil {
		log.Fatal("listen tcp error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept tcp error:", err)
		}
		go handleConn(conn, remote)
	}
}
