package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"

	"os"
	"time"

	"github.com/NoobforAl/DnsLab/core"
	log "github.com/sirupsen/logrus"
)

// * get ip/host/port from user
func getPortHostIp() (ipHost string, port uint16) {
	log.Warn("if you inter empty port or ip, ")
	log.Warn("program set default yor ip port.")

	fmt.Println("Enter Host/Ip:")
	fmt.Scanln(&ipHost)

	fmt.Println("Enter Port:")
	fmt.Scanln(&port)

	return
}

// * check enter query type is valid
func checkQueryType(q *string) error {
	switch *q {
	case "1", "2", "5", "6", "16":
		return nil
	default:
		return fmt.Errorf("not found query type")
	}
}

func runCommands(
	showIP *bool,
	token *string,
	pingIP, openPort *bool,
	dnsLookup *string,
	reverseLookup, updateIP, upApp *bool,
	timeUpdate *int,
) int {
	app := core.BaseConf{Token: *token}

	if err := checkQueryType(dnsLookup); *dnsLookup != "" && err != nil {
		log.Error(err)
		return 1

	} else if *dnsLookup != "" {
		ipHost, port := getPortHostIp()
		app.SetUpIp(ipHost, "", port)

		res, err := app.DnsLookup(*dnsLookup)
		if err != nil {
			log.Error(err)
			return 1
		}

		b, err := io.ReadAll(res.Body)
		if err != nil {
			res.Body.Close()
			log.Error(err)
			return 1
		}

		var data []map[string]interface{}
		err = json.Unmarshal(b, &data)
		if err != nil {
			res.Body.Close()
			log.Error(err)
			return 1
		}

		res.Body.Close()

		for _, d := range data {
			for k, v := range d {
				log.Printf("%s: %v", k, v)
			}
		}
	}

	if *pingIP {
		ipHost, port := getPortHostIp()
		app.SetUpIp(ipHost, "", port)

		v, err := app.Ping()
		if err != nil {
			log.Error(err)
			return 1
		}
		log.Printf("BufferSize: %d", v.BufferSize)
		log.Printf("Ip: %s", v.Ip)
		log.Printf("Ttl: %d", v.Ttl)
		log.Printf("Time: %d", v.Time)
		log.Printf("Success: %v", v.Success)
	}

	if *openPort {
		ipHost, port := getPortHostIp()
		app.SetUpIp(ipHost, "", port)

		v, err := app.OpenPort()
		if err != nil {
			log.Error(err)
			return 1
		}
		log.Printf("Port is open: %v\n", v)
	}

	if *reverseLookup {
		ipHost, port := getPortHostIp()
		app.SetUpIp(ipHost, "", port)

		v, err := app.ReverseLookup()
		if err != nil {
			log.Error(err)
			return 1
		}
		log.Printf("DNS: %s", v)
	}

	if *showIP {
		if err := app.CheckIP(); err != nil {
			log.Error(err)
			return 1
		}

		log.Println("Ipv4: ", app.Ip)
		log.Println("Ipv6: ", app.Ipv6)
	}

	if *updateIP {
		v, err := app.UpdateIp()
		if err != nil {
			log.Error(err)
			return 1
		}
		log.Printf("IP updated: %v", v)
	}

	if *upApp {
		for {
			v, err := app.UpdateIp()
			if err != nil {
				log.Error(err)
				return 1
			}
			log.Printf("IP updated: %v", v)
			time.Sleep(time.Duration(*timeUpdate) * time.Minute)
		}
	}

	return 0
}

func Run() int {
	showIP := flag.Bool("ip", false, "See Your Ip")
	token := flag.String("t", "", "Set your token use -t your token")
	pingIP := flag.Bool("pi", false, "Ping your IP")
	openPort := flag.Bool("op", false, "Open Port Checker")
	dnsLookup := flag.String("dl", "", "Dns Lookup ues -dl query type")
	reverseLookup := flag.Bool("rl", false, "Reverse Lookup")
	updateIP := flag.Bool("uip", false, "Update your ip with token! use this command with -t <your token>")
	upApp := flag.Bool("up", false, "every 3m or any time check your ip!")
	timeUpdate := flag.Int("ts", 3, "time sleep for check ip, default is 3m")
	flag.Parse()

	// * check program run with args
	// * if not show help and exit
	if len(os.Args) <= 1 {
		flag.PrintDefaults()
		return 0
	}

	return runCommands(
		showIP, token,
		pingIP, openPort,
		dnsLookup, reverseLookup,
		updateIP, upApp, timeUpdate,
	)
}
