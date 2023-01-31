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

func getQuery() (s string, err error) {
	log.Warn("Enter Query :")
	fmt.Scanln(&s)
	err = getQueryValue(s)
	return
}

func getQueryValue(q string) error {
	switch q {
	case "1", "2", "5", "6", "16":
		return nil
	default:
		return fmt.Errorf("not found query type")
	}
}
