package http

import (
	"github.com/ian-kent/go-log/log"
	gotcha "github.com/ian-kent/gotcha/app"
)

func Listen(httpBindAddr string, Asset func(string) ([]byte, error), exitCh chan int, registerCallback func(*gotcha.App)) {
	log.Info("[HTTP] Binding to address: %s", httpBindAddr)

	var app = gotcha.Create(Asset)
	app.Config.Listen = httpBindAddr

	registerCallback(app)

	app.Start()

	<-make(chan int)
	exitCh <- 1
}
