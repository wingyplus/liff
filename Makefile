VERSION := v1.1.0

build-windows:
	mkdir -p build/windows
	CGO_ENABLED=0 GOOS=windows GOARC=amd64 go build -o ./build/windows/liff.exe

build-linux:
	mkdir -p build/linux
	CGO_ENABLED=0 GOOS=linux GOARC=amd64 go build -o ./build/linux/liff

build-macos:
	mkdir -p build/macos
	CGO_ENABLED=0 GOOS=darwin GOARC=amd64 go build -o ./build/macos/liff

pkg-windows: build-windows
	zip -j ./build/liff-windows-amd64-$(VERSION).zip ./build/windows/liff.exe

pkg-linux: build-linux
	tar czvf ./build/liff-linux-amd64-$(VERSION).tar.gz ./build/linux/liff

pkg-macos: build-macos
	tar czvf ./build/liff-macos-amd64-$(VERSION).tar.gz ./build/macos/liff

pkg: pkg-macos pkg-linux pkg-windows
.PHONY: package pkg-macos pkg-linux pkg-windows build-macos build-linux build-windows
