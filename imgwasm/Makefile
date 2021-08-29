SHELL := /bin/bash
.PHONY: check

.SILENT: build getlasttag buildzip


build:
	$(info +Компиляция Wasm)
	GOOS=js GOARCH=wasm go build -ldflags "-s -w" -o ../httpsrv/www/main.wasm main/main.go
#go build -ldflags "-s -w" -o ./bin/main/canvas cmd/main/main.go

buildserver: 
	go build -o ../httpsrv/www/server  ../httpsrv/main/main.go

run: build buildserver
	$(info +Запуск)
	../httpsrv/www/server

buildtiny:
	$(info +Компиляция Wasm)
	GOPATH=/usr/local/go GOROOT=/usr/local/go GOOS=js GOARCH=wasm tinygo build -o ../httpsrv/www/main.wasm main/main.go
#go build -v -ldflags="-X 'main.Version=v1.0.0' -X 'app/build.User=$(id -u -n)' -X 'app/build.Time=$(date)'"