package core

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func getIpPort(getPort bool, text string) (ip string, port uint16) {
	log.Warn(text)
	fmt.Scanln(&ip)

	if getPort {
		log.Print("Enter port or enter:")
		fmt.Scanln(&port)
	}
	return
}
