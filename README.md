# DnsLab-cli

<div dir="rtl">

یک cli برای وب سایت [dnslab](dnslab.link) به منظور راحتی در دست رسی به api وب سایت.

پروژه اصلی [لینک](https://github.com/AkbarAsghari/DNSLab-WebSite)

## راهنمای نصب

برای نصب ابزار نیاز به نصب Go
دارید برای نصب به این [link](https://go.dev/doc/install) .مراجعه کنید

بعد از نصب این کامند را اجرا کنید.

<div dir="ltr">

> go install github.com/NoobforAl/DnsLab@latest

</div>

بعد از نصب Dnslab
چک کنید که به درستی ابزار نصب شده است.

<div dir="ltr">

```
$ dnslab
  -dl string
        Dns Lookup ues -dl query type
  -ip
        See Your Ip
  -op
        Open Port Checker
  -pi
        Ping your IP
  -rl
        Reverse Lookup
  -t string
        Set your token use -t your token
  -ts int
        time sleep for check ip, default is 3m (default 3)
  -uip
        Update your ip with token! use this command with -t <your token>
  -up
        every 3m or any time check your ip!

```

</div>

در صورتی که تمایلی به نصب Go
ندارید می توانید باینری فایل را برای خود دانلود کنید.
[releases](https://github.com/NoobforAl/DnsLab/releases)

توجه: برای تمامی سیستم عامل ها تعبیه نشده است
و ممکن است برای برخی از معماری پردازنده ها درست کار نکند.

# راهنمای CLI

<div dir="ltr">

| کامند |                                         توضیحات                                          |
| ----- | :--------------------------------------------------------------------------------------: |
| -h    |                                       دیدن راهنما                                        |
| -dl "query type"  |                                  دیدن ای پی با دی ان اس                                  |
| -ip   |                  دیدن ای پی به صورت پیش فرض ای پی شما نمایش داده می شود                  |
| -op   |                                   چک کردن پورت ها باز                                    |
| -pi   |                        پینگ کردن ای پی خود یا مقداری که وارد شده                         |
| -rl   |                                  دیدن دی ان اس با ای پی                                  |
| -t  "your token"  |                توکنی که از وب سایت دریافت کرده اید با این دستور وارد کنید                |
| -ts "minute"  |                             زمان مرتب چک کردن ای پی به دقیقه                              |
| -uip  |                   اپدیت کردن ای پی به شرطی که توکنی را وارد کرده باشید                   |
| -up   | چک کردن مداوم ای پی بسته به زمان تنظیم شده به صورت پیش فرض بر هر 3 دقیقه یک بار تنظیم است |

</div>

# Example

آپدیت کردن IP

<div dir="ltr">

```bash
$ dnslab -uip -t your token
 IP updated:  true
 IP updated:  true
```

</div>

</div>
