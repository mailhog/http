package http

import (
	"net/http"

	"github.com/gorilla/pat"
	"github.com/ian-kent/go-log/log"
)

// Listen binds to httpBindAddr
func Listen(httpBindAddr string, Asset func(string) ([]byte, error), exitCh chan int, registerCallback func(http.Handler)) {
	log.Info("[HTTP] Binding to address: %s", httpBindAddr)

	pat := pat.New()
	registerCallback(pat)

	err := http.ListenAndServe(httpBindAddr, pat)
	if err != nil {
		log.Fatalf("[HTTP] Error binding to address %s: %s", httpBindAddr, err)
	}
}
