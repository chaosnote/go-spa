# 專案( Stock )前端

## Build
windows output exe
1. go build -o ./dist/bro.exe -race main.go

windows output linux
1. powershell
1. $env:GOOS = "linux"
1. go build -o ./dist/Main main.go
