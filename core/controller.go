package core

import "fmt"

func getIpPort(getPort bool, text string) (ip string, port uint16) {
	fmt.Print(text)
	fmt.Scanln(&ip)

	if getPort {
		fmt.Print("Enter port or enter:")
		fmt.Scanln(&port)
	}
	return
}
