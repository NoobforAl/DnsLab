package main

import (
	"flag"

	"time"
)

var (
	pingIP = flag.Bool("pi", false, "Ping your IP")
	showIP = flag.Bool("show-ip", false, "See Your Ip")

	debug = flag.Bool("d", false, "run debug mode")

	openPort      = flag.Bool("op", false, "Open Port Checker")
	reverseLookup = flag.Bool("rl", false, "Reverse Lookup")

	upApp    = flag.Bool("check-ip", false, "every 3m or any time check your ip!")
	updateIP = flag.Bool("update-ip", false, "Update your ip with token! use this command with -t <your token>")

	ipOrHost  = flag.String("addr", "", "Set your ip or host")
	dnsLookup = flag.String("dl", "", "Dns Lookup ues -dl query type")
	token     = flag.String("token", "", "Set your token use -t your token")

	timeUpdate = flag.Duration("time-check", time.Minute*3, "time sleep for check ip, default is 3m")
	retryTime  = flag.Duration("retry-time", time.Second*3, "time sleep for request again if get error, default is 3s")

	port       = flag.Uint("port", 80, "set your port for request")
	retryCount = flag.Uint("retry-count", 3, "how many time try for request again, default is 3  times")
)
