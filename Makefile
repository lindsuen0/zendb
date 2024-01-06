APP := zendb
VERSION := 1.0.0
DIR := bin

.PHONY: all clean build linux freebsd

all: linux freebsd

clean:
	@if [ -d ${DIR} ]; then rm -rf ${DIR}/*; else exit 0; fi

build:
	go build -o ${DIR}/${APP} main.go

linux:
	@# linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o ${DIR}/${APP} main.go
	@cd ${DIR}/ && tar -zcf ${APP}-${VERSION}-linux_amd64.tar.gz ${APP} && rm -rf ${APP} && cd ../
	@# linux-arm64:
	GOOS=linux GOARCH=arm64 go build -o ${DIR}/${APP} main.go
	@cd ${DIR}/ && tar -zcf ${APP}-${VERSION}-linux_arm64.tar.gz ${APP} && rm -rf ${APP} && cd ../

freebsd:
	@# freebsd-amd64:
	GOOS=freebsd GOARCH=amd64 go build -o ${DIR}/${APP} main.go
	@cd ${DIR}/ && tar -jcf ${APP}-${VERSION}-freebsd_amd64.tar.bz2 ${APP} && rm -rf ${APP} && cd ../
	@# freebsd-arm64:
	GOOS=freebsd GOARCH=arm64 go build -o ${DIR}/${APP} main.go
	@cd ${DIR}/ && tar -jcf ${APP}-${VERSION}-freebsd_arm64.tar.bz2 ${APP} && rm -rf ${APP} && cd ../
