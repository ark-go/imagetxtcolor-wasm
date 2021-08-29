SHELL := /bin/bash
.PHONY: check

.SILENT: build getlasttag buildzip


build:
	$(info +Компиляция Server)
	go build -o www/server  main/main.go

run: build
	$(info +Запуск)
	./www/server
