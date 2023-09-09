package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/NoobforAl/DnsLab/src/contract"
	"github.com/NoobforAl/DnsLab/src/core"
	"github.com/sirupsen/logrus"
)

var (
	app contract.DnslabAPi
	log contract.Logger
)

func init() {
	// * check program run with args
	// * if not show help and exit
	if len(os.Args) <= 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()

	logLevel := logrus.InfoLevel
	if *debug {
		logLevel = logrus.DebugLevel
	}

	log = &logrus.Logger{
		Out:          os.Stderr,
		Formatter:    new(logrus.TextFormatter),
		Hooks:        make(logrus.LevelHooks),
		Level:        logLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}

	app = core.New(*token, log, *retryTime, *retryCount)
}

func main() {

	// handle panic
	// and ctrl+c key
	defer recoverPanic()
	go handelKeyInput()

	var err error

	if err = checkQueryType(dnsLookup); *dnsLookup != "" && err != nil {
		log.Error(err)
	} else if *dnsLookup != "" {
		data, err := app.DnsLookup(*ipOrHost, *dnsLookup)
		if err != nil {
			log.Error(err)
		}
		printData(data)
	}

	if *pingIP {
		data, err := app.Ping(*ipOrHost)
		if err != nil {
			log.Error(err)
		}
		printData(data)
	}

	if *openPort {
		data, err := app.OpenPort(*ipOrHost, uint16(*port))
		if err != nil {
			log.Error(err)
		}
		printData(data)
	}

	if *reverseLookup {
		data, err := app.ReverseLookup(*ipOrHost)
		if err != nil {
			log.Error(err)
		}
		printData(fmt.Sprintf("DNS: %s", data))
	}

	if *showIP {
		data, err := app.CheckIP()
		if err != nil {
			log.Error(err)
		}

		printData(fmt.Sprintf("DNS: %v", data))
	}

	if *updateIP {
		v, err := app.UpdateIp()
		if err != nil {
			log.Error(err)
		}
		printData(fmt.Sprintf("IP updated: %v", v))
	}

	if *upApp {
		for {
			v, err := app.UpdateIp()
			if err != nil {
				log.Fatal(err)
			}
			printData(fmt.Sprintf("IP updated: %v", v))
			<-time.After(*timeUpdate)
		}
	}
}
