# 專案( Stock )前端

## github tree
project
    1. svr
    1. https://github.com/chaosnote/kernel
    1. https://github.com/chaosnote/build

## Build
windows output exe
1. go build -o ./dist/bro.exe -race main.go

windows output linux
1. powershell
1. $env:GOOS = "linux"
1. go build -o ./dist/Main main.go
