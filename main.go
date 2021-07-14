package main

import (
	"github.com/chaosnote/go-build/log"
	"github.com/chaosnote/go-kernel/net/server"
	"github.com/chaosnote/go-kernel/schedule"

	"svr/web"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	log.Set("dist", "bro", true)

	port := "9090"

	log.Console("host", zap.String("url", "http://localhost:"+port))

	r := mux.NewRouter()
	web.Register(r)
	server.New(":"+port, r)

	log.Console("ready")

	schedule.Run() // always

}
