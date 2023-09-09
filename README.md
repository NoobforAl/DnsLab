# DnsLab-cli

<div dir="rtl">

یک cli برای وب سایت [dnslab](https://dnslab.link) به منظور راحتی در دست رسی به api وب سایت.

DnsLab Project [Link](https://github.com/AkbarAsghari/DNSLab-WebSite)

## راهنمای نصب

برای استفاده از ابزار به 
[releases](https://github.com/NoobforAl/DnsLab/releases) مراجعه کنید و باینری فایل را دانلود و اجرا کنید.

توجه: برای تمامی سیستم عامل ها تعبیه نشده است
و ممکن است برای برخی از معماری پردازنده ها درست کار نکند.

# راهنمای CLI

<div dir="ltr">

```txt
$ dnslab
  -addr string
        Set your ip or host
  -check-ip
        every 3m or any time check your ip!
  -d    run debug mode
  -dl string
        Dns Lookup ues -dl query type
  -op
        Open Port Checker
  -pi
        Ping your IP
  -port uint
        set your port for request (default 80)
  -retry-count uint
        how many time try for request again, default is 3  times (default 3)
  -retry-time duration
        time sleep for request again if get error, default is 3s (default 3s)
  -rl
        Reverse Lookup
  -show-ip
        See Your Ip
  -time-check duration
        time sleep for check ip, default is 3m (default 3m0s)
  -token string
        Set your token use -t your token
  -update-ip
        Update your ip with token! use this command with -t <your token>

```

</div>
