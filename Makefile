# goEncrypt Makefile

default:
	go build goEncrypt.go

install:
	go build goEncrypt.go
	cp goEncrypt /usr/bin


clean:
	rm goEncrypt
