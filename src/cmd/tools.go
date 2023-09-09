package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

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
