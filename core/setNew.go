package core

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
