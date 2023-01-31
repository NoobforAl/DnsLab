package core

type BaseConf struct {
	Token string
	Ip    string
	Ipv6  string
	Port  uint16
}

func (c *BaseConf) CheckIP() error {
	return c.checkIP()
}

func (c *BaseConf) Ping() (*ipOrHostPing, error) {
	defer c.setUpIp(c.Ip, c.Ipv6, 0)
	c.setUpEnterIp(false, ENTER_IP)
	return c.ping()
}

func (c *BaseConf) OpenPort() (bool, error) {
	defer c.setUpIp(c.Ip, c.Ipv6, 0)
	c.setUpEnterIp(true, ENTER_IP)
	return c.openPort()
}

func (c *BaseConf) DnsLookup() (data qTypes, err error) {
	defer c.setUpIp(c.Ip, c.Ipv6, 0)
	c.setUpEnterIp(false, ENTER_DNS)
	q, err := getQuery()
	if err != nil {
		return
	}
	return c.dnsLookup(q)
}

func (c *BaseConf) ReverseLookup() (string, error) {
	defer c.setUpIp(c.Ip, c.Ipv6, 0)
	c.setUpEnterIp(false, ENTER_DNS)
	return c.reverseLookup()
}

func (c *BaseConf) UpdateIp() (bool, error) {
	defer c.setUpIp(c.Ip, c.Ipv6, 0)
	return c.updateIp()
}
