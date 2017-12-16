package AnotherSteamCommunityFix

import "net/http"

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
	http.ListenAndServe(":80", http.HandlerFunc(redirect))
}
