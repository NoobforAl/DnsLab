package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/NoobforAl/DnsLab/core"
	log "github.com/sirupsen/logrus"
)

var app core.BaseConf = core.BaseConf{}

func init() {
	if len(os.Args) >= 1 {
		if err := app.CheckIP(); err != nil {
			log.Fatal(err.Error())
		}
	}
}

func main() {
	showIP := flag.Bool("ip", false, "See Your Ip")
	token := flag.String("t", "", "Set your token")
	pingIP := flag.Bool("pi", false, "Ping your IP")
	openPort := flag.Bool("op", false, "Open Port Checker")
	dnsLookup := flag.Bool("dl", false, "Dns Lookup")
	reverseLookup := flag.Bool("rl", false, "Reverse Lookup")
	updateIP := flag.Bool("uip", false, "Update your ip with token!")
	upApp := flag.Bool("up", false, "every 3h or any time check your ip!")
	timeUpdate := flag.Int("ts", 3, "time sleep for check ip")
	flag.Parse()

	app.Token = *token

	if len(os.Args) <= 1 {
		flag.PrintDefaults()
		return
	}

	if app.Token != "" {
		v, err := app.UpdateIp()
		app.Show(err, fmt.Sprintf("IP updated: %v", v))
	}

	if *pingIP {
		v, err := app.Ping()
		app.Show(err,
			fmt.Sprintf("BufferSize: %d", v.BufferSize),
			fmt.Sprintf("Ip: %s", v.Ip),
			fmt.Sprintf("Ttl: %d", v.Ttl),
			fmt.Sprintf("Time: %d", v.Time),
			fmt.Sprintf("Success: %v", v.Success),
		)
	}

	if *openPort {
		v, err := app.OpenPort()
		app.Show(err, fmt.Sprintf("Port is open: %v", v))
	}

	if *dnsLookup {
		v, err := app.DnsLookup()
		app.Show(err, fmt.Sprintf("Ip: %s", v))
	}

	if *reverseLookup {
		v, err := app.ReverseLookup()
		app.Show(err, fmt.Sprintf("DNS: %s", v))
	}

	if *updateIP {
		v, err := app.UpdateIp()
		app.Show(err, fmt.Sprintf("IP updated: %v", v))
	}

	if *showIP {
		log.Println("Ipv4: ", app.Ip)
		log.Println("Ipv6: ", app.Ipv6)
	}

	if *upApp {
		for {
			v, err := app.UpdateIp()
			app.Show(err, fmt.Sprintf("IP updated: %v", v))
			time.Sleep(time.Duration(*timeUpdate) * time.Hour)
		}
	}
}
