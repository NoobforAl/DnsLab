package core

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

const (
	API           string = "https://api.dnslab.ir/"
	IP            string = API + "IP/"
	PING          string = API + "IP/Ping?hostOrIPAddress=%v"
	OPEN          string = API + "IP/IsIPAndPortOpen?hostOrIPAddress=%v&port=%v"
	DNSLOOKUP     string = API + "IP/DNSLookup?hostOrIPAddress=%v"
	REVERSELookup string = API + "IP/ReverseLookup?IPAddress=%v"
	DNSREF        string = API + "DNS/UDIBT?k=%v"

	ENTER_IP  string = "if enter, your ip checked!\nEnter ip or enter: "
	ENTER_DNS string = "Enter dns: "
)

type BaseConf struct {
	Token string
	Ip    string
	Ipv6  string
	Port  uint16
}

func (c *BaseConf) Show(text, v any, err error) {
	if err != nil {
		log.Panicln(err.Error())
	} else {
		log.Println(text, v)
	}
}

func (c *BaseConf) setUpIp(IPv4, IPv6 string, port uint16) {
	c.Ip = IPv4
	c.Ipv6 = IPv6

	c.Port = port
	if port == 0 {
		c.Port = 80
	}
}

func (c *BaseConf) setUpEnterIp(getPort bool, text string) {
	ip, port := getIpPort(getPort, text)
	if ip != "" {
		c.Ip = ip
	}

	if port != 0 {
		c.Port = port
	}
}

func (c *BaseConf) CheckIP() error {
	res, err := request(IP)
	if err != nil {
		return err
	}

	r, err := decodeBodyJson(res)
	if err != nil {
		return err
	}

	c.setUpIp(r.IPv4, r.IPv6, 0)
	return err
}

func (c *BaseConf) Ping() (bool, error) {
	defer c.setUpIp(c.Ip, c.Ipv6, 0)

	c.setUpEnterIp(false, ENTER_IP)
	url := fmt.Sprintf(PING, c.Ip)
	res, err := request(url)
	if err != nil {
		return false, err
	}
	return boolPars(res)
}

func (c *BaseConf) OpenPort() (bool, error) {
	defer c.setUpIp(c.Ip, c.Ipv6, 0)

	c.setUpEnterIp(true, ENTER_IP)
	url := fmt.Sprintf(OPEN, c.Ip, c.Port)
	res, err := request(url)
	if err != nil {
		return false, err
	}
	return boolPars(res)
}

func (c *BaseConf) DnsLookup() (string, error) {
	defer c.setUpIp(c.Ip, c.Ipv6, 0)

	c.setUpEnterIp(false, ENTER_DNS)
	url := fmt.Sprintf(DNSLOOKUP, c.Ip)
	res, err := request(url)
	if err != nil {
		return "", err
	}
	return bodyToString(res)
}

func (c *BaseConf) ReverseLookup() (string, error) {
	defer c.setUpIp(c.Ip, c.Ipv6, 0)

	c.setUpEnterIp(false, ENTER_DNS)
	url := fmt.Sprintf(DNSLOOKUP, c.Ip)
	res, err := request(url)
	if err != nil {
		return "", err
	}
	return bodyToString(res)
}

func (c *BaseConf) UpdateIp() (bool, error) {
	defer c.setUpIp(c.Ip, c.Ipv6, 0)

	url := fmt.Sprintf(DNSREF, c.Token)
	res, err := request(url)
	if err != nil {
		return false, err
	}
	return boolPars(res)
}
