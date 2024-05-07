package http

import (
	"fmt"
	"html"
	"log"
	net_http "net/http"
)

func HTTPServe(addr string) {
	net_http.HandleFunc("/bar", func(w net_http.ResponseWriter, r *net_http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(net_http.ListenAndServe(addr, nil))
}
