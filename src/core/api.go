package core

/*
* all url and path for request to api
* %s setup in program
 */
const (
	API           string = "https://api.dnslab.link/"
	IP            string = API + "IP/"
	PING          string = API + "IP/Ping?hostOrIPAddress=%s"
	OPEN          string = API + "IP/IsIPAndPortOpen?hostOrIPAddress=%s&port=%d"
	DNSLOOKUP     string = API + "DNSLookUp/Query?Query=%s&queryType=%s"
	REVERSELookup string = API + "DNSLookUp/QueryReverse?IPAddress=%s"
	DNSREF        string = API + "DNS/UDIBT?k=%s"
)
