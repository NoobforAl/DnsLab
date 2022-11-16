package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/NoobforAl/DnsLab/core"
)

var app core.BaseConf

func init() {
	app = core.BaseConf{}
	if err := app.CheckIP(); err != nil {
		fmt.Println(err.Error())
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
	fmt.Println(" Ipv4: ", app.Ip)
	fmt.Println(" Ipv6: ", app.Ipv6)

	if app.Token != "" {
		v, err := app.UpdateIp()
		app.Show(" IP updated: ", v, err)
	}

	if *showIP {
		fmt.Println(" Ipv4: ", app.Ip)
		fmt.Println(" Ipv6: ", app.Ipv6)
	}

	if *pingIP {
		v, err := app.Ping()
		app.Show(" IP is up(PING): ", v, err)
	}

	if *openPort {
		v, err := app.OpenPort()
		app.Show(" Port is opne: ", v, err)
	}

	if *dnsLookup {
		v, err := app.DnsLookup()
		app.Show(" Ip: ", v, err)
	}

	if *reverseLookup {
		v, err := app.ReverseLookup()
		app.Show(" DNS: ", v, err)
	}

	if *updateIP {
		v, err := app.UpdateIp()
		app.Show(" IP updated: ", v, err)
	}

	if *upApp {
		for {
			v, err := app.UpdateIp()
			app.Show(" IP upadated: ", v, err)
			time.Sleep(time.Duration(*timeUpdate) * time.Hour)
		}
	}
	time.Sleep(10 * time.Second)
}
