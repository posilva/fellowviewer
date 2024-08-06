
.PHONY: build windows linux macos macos_arm64

build: macos_arm64 windows linux macos

windows:
	GOOS=windows GOARCH=amd64 go build -o bin/windows/fellowviewer.exe main.go

linux:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/fellowviewer main.go

macos:
	GOOS=darwin GOARCH=amd64 go build -o bin/macos/fellowviewer main.go
macos_arm64:
	GOOS=darwin GOARCH=arm64 go build -o bin/macos_arm64/fellowviewer main.go
