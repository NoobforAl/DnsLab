// ! wrong: all query type in program not used!
// ! this for see what's response
package core

/*
* and setup value response to config struct
 */
type ipInfo struct {
	IPv4 string `json:"iPv4"`
	IPv6 string `json:"iPv6"`
}

/*
* use for ping your ip or other ip
 */
type ipOrHostPing struct {
	BufferSize int    `json:"bufferSize"`
	Ip         string `json:"ip"`
	Ttl        int    `json:"ttl"`
	Time       int    `json:"time"`
	Success    bool   `json:"success"`
}

/*
* this interface use for decoder json
* and some returned value then we don't know type
 */
type response interface {
	ipOrHostPing | ipInfo
}

/*
* this struct is base info for all query type
* then we request
* and model for all query type struct
* ( DnsLookUpQT1, DnsLookUpQT2,
*	DnsLookUpQT5, DnsLookUpQT6, DnsLookUpQT16)
 */
type dnsLookUpQT struct {
	RecordType string `json:"recordType"`
	DomainName string `json:"domainName"`
	Ttl        int    `json:"ttl"`
}

// * model for query type 1
type DnsLookUpQT1 struct {
	dnsLookUpQT

	Address string `json:"address"`
}

// * model for query type 2
type DnsLookUpQT2 struct {
	dnsLookUpQT

	NsdName string `json:"nsdName"`
}

// * model for query type 5
type DnsLookUpQT5 struct {
	dnsLookUpQT

	CanonicalName string `json:"canonicalName"`
}

// * model for query type 6
type DnsLookUpQT6 struct {
	dnsLookUpQT

	Expire  int    `json:"expire"`
	Minimum int    `json:"minimum"`
	MName   string `json:"mName"`
	Refresh int    `json:"refresh"`
	Retry   int    `json:"retry"`
	RName   string `json:"rName"`
	Serial  int    `json:"serial"`
}

// * model for query type 16
type DnsLookUpQT16 struct {
	dnsLookUpQT

	Exchange   string `json:"exchange"`
	Preference int    `json:"preference"`
}
