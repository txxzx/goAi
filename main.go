package main

import (
	td "github.com/swxctx/malatd"
	"github.com/txxzx/goAi/api"
)

func main() {
	// Gen Time: 2024-04-16 16:07:44
	srv := td.NewServer(cfg.SrvConfig)
	api.Route(srv, "/go_ai")
	srv.Run()
}
