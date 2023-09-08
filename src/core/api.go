package core

/*
* all url and path for request to api
* %s set in program
 */
const (
	API           = "https://api.dnslab.link/"
	IP            = API + "IP/"
	PING          = API + "IP/Ping?hostOrIPAddress=%s"
	OPEN          = API + "IP/IsIPAndPortOpen?hostOrIPAddress=%s&port=%d"
	DNS_LOOKUP    = API + "DNSLookUp/Query?Query=%s&queryType=%s"
	REVERSELookup = API + "DNSLookUp/QueryReverse?IPAddress=%s"
	DNS_REF       = API + "DNS/UDIBT?k=%s"
)
