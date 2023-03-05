package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/NoobforAl/DnsLab/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	defer recoverPanic()
	go handelKeyInput()
	os.Exit(cmd.Run())
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
