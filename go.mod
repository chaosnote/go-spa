module svr

go 1.14

replace kernel => ../kernel

replace build => ../build

require (
	build v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
	go.uber.org/zap v1.15.0
	kernel v0.0.0-00010101000000-000000000000
)
