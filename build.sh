#!/bin/bash

VERSION="v0.6.0"

WINDOWS="./build/windows-x86-64bit-$VERSION"
LINUX="./build/linux-x86-64bit-$VERSION"
MAC_OS_ARM="./build/macOs-arm-64bit-$VERSION"
MAC_OS="./build/macOs-x86-64bit-$VERSION"
MAIN_FILE="src/cmd/main.go"

go mod tidy

export GOOS=windows
go build -o "$WINDOWS/dnslab.exe" $MAIN_FILE
zip -r $WINDOWS.zip $WINDOWS/


export GOOS=linux
go build -o "$LINUX/dnslab" $MAIN_FILE
zip -r $LINUX.zip $LINUX/


export GOOS=darwin
go build -o "$MAC_OS/dnslab" $MAIN_FILE
zip -r $MAC_OS.zip $MAC_OS/


export GOOS=darwin
export GOARCH=arm64
go build -o "$MAC_OS_ARM/dnslab" $MAIN_FILE
zip -r $MAC_OS_ARM.zip $MAC_OS_ARM/