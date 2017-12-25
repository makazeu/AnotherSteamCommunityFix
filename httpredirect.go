package AnotherSteamCommunityFix

import (
	"net/http"
	"log"
	"net"
)

var statusCode int

func redirect(w http.ResponseWriter, req *http.Request) {
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	http.Redirect(w, req, target, statusCode)
}

func StartServingHTTPRedirect(httpCode int) {
	statusCode = httpCode
	log.Fatal(ListenAndServe(":80", http.HandlerFunc(redirect)))
}

type Server struct {
	http.Server
}

func ListenAndServe(addr string, handler http.Handler) error {
	server := &Server{http.Server{Addr:addr, Handler:handler}}
	return server.ListenAndServe()
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