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
	Success    bool   `json:"success"`
}

type dnsLookUpQT1 struct {
	RecordType string `json:"recordType"`
	DomainName string `json:"domainName"`
	Ttl        int    `json:"ttl"`
	Address    string `json:"address"`
}

type dnsLookUpQT2 struct {
	RecordType string `json:"recordType"`
	DomainName string `json:"domainName"`
	Ttl        int    `json:"ttl"`
	NsdName    string `json:"nsdName"`
}

type dnsLookUpQT5 struct {
	RecordType    string `json:"recordType"`
	DomainName    string `json:"domainName"`
	Ttl           int    `json:"ttl"`
	CanonicalName string `json:"canonicalName"`
}

type dnsLookUpQT6 struct {
	RecordType string `json:"recordType"`
	DomainName string `json:"domainName"`
	Ttl        int    `json:"ttl"`
	Expire     int    `json:"expire"`
	Minimum    int    `json:"minimum"`
	MName      string `json:"mName"`
	Refresh    int    `json:"refresh"`
	Retry      int    `json:"retry"`
	RName      string `json:"rName"`
	Serial     int    `json:"serial"`
}

type dnsLookUpQT16 struct {
	RecordType string `json:"recordType"`
	DomainName string `json:"domainName"`
	Ttl        int    `json:"ttl"`
	Exchange   string `json:"exchange"`
	Preference int    `json:"preference"`
}

type qTypes struct {
	Q1  []dnsLookUpQT1
	Q2  []dnsLookUpQT2
	Q5  []dnsLookUpQT5
	Q6  []dnsLookUpQT6
	Q16 []dnsLookUpQT16
}

type response interface {
	ipOrHostPing | ipInfo |
		[]dnsLookUpQT1 | []dnsLookUpQT2 |
		[]dnsLookUpQT5 | []dnsLookUpQT6 |
		[]dnsLookUpQT16
}
