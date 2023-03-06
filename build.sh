#!/bin/bash

VERSION="v0.5.1"

WINDOWS="./build/windows-x86-64bit-$VERSION"
LINUX="./build/linux-x86-64bit-$VERSION"
MAC_OS_ARM="./build/macOs-arm-64bit-$VERSION"
MAC_OS="./build/macOs-x86-64bit-$VERSION"

go mod tidy

export GOOS=windows
go build -o "$WINDOWS/dnslab.exe" main.go
zip -r $WINDOWS.zip $WINDOWS/


export GOOS=linux
go build -o "$LINUX/dnslab" main.go
zip -r $LINUX.zip $LINUX/


export GOOS=darwin
go build -o "$MAC_OS/dnslab" main.go
zip -r $MAC_OS.zip $MAC_OS/


export GOOS=darwin
export GOARCH=arm64
go build -o "$MAC_OS_ARM/dnslab" main.go
zip -r $MAC_OS_ARM.zip $MAC_OS_ARM/