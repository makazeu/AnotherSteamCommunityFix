package AnotherSteamCommunityFix

import (
	"net"
	"net/http"
)

func redirectHTTPS(w http.ResponseWriter, req *http.Request) {
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	http.Redirect(w, req, target, http.StatusFound)
}

func StartServingHTTPSRedirect() {
	server := &Server{
		http.Server{
			Addr:    ":80",
			Handler: http.HandlerFunc(redirectHTTPS),
		},
	}
	fatalLogger.Fatal(server.ListenAndServe())
}

type Server struct {
	http.Server
}

func (srv *Server) ListenAndServe() error {
	addr := srv.Addr
	if addr == "" {
		addr = ":http"
	}
	ln, err := net.Listen("tcp4", addr)
	if err != nil {
		return err
	}
	return srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
}

type tcpKeepAliveListener struct {
	*net.TCPListener
}
