package core

type ipInfo struct {
	IPv4 string `json:"iPv4"`
	IPv6 string `json:"iPv6"`
}

type ipOrHostPing struct {
	BufferSize int    `json:"bufferSize"`
	Ip         string `json:"ip"`
	Ttl        int    `json:"ttl"`
	Time       int    `json:"time"`
	Success    string `json:"success"`
}

type response interface {
	ipOrHostPing | ipInfo
}
