package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NoobforAl/DnsLab/src/contract"
	"github.com/NoobforAl/DnsLab/src/core"
	"github.com/sirupsen/logrus"
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

func printData(m any) {
	switch m := m.(type) {
	case string:
		log.Infof("%s", m)

	case bool:
		log.Infof("%b", m)

	case map[string]any:
		for k, v := range m {
			log.Infof("%s: %v", k, v)
		}

	case []map[string]any:
		for _, v := range m {
			log.Info("------------------")
			printData(v)
			log.Info("------------------")
		}
	}
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

/*
* Handel ctrl + c keyPut
 */
func handelKeyInput() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGALRM)

	// * wait for get data from chanel
	<-c

	log.Println("Exit Program....")
	log.Println("Good bye!")
	os.Exit(0)
}

/*
* recover panic code
 */
func recoverPanic() {
	if err := recover(); err != nil {
		log.Warn(err)
		os.Exit(1)
	}
}
