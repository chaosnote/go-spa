package main

import (
	"kernel/net/server"
	"kernel/schedule"

	"svr/tool/log"
	"svr/web"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	log.Set("bro", true)

	port := "9090"

	log.Console("host", zap.String("url", "http://localhost:"+port))

	r := mux.NewRouter()
	web.Register(r)
	server.New(":"+port, r)

	log.Console("ready")

	schedule.Run() // always

}
