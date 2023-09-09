package contract

type DnslabAPi interface {
	UpdateIp() (bool, error)
	CheckIP() (map[string]any, error)
	Ping(ipOrHost string) (map[string]any, error)

	ReverseLookup(ip string) (string, error)
	OpenPort(ipOrHost string, prot uint16) (bool, error)
	DnsLookup(ipOrHost, q string) (map[string]any, error)
}
