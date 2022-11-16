# DnsLab-cli

<div dir="rtl">

یک cli برای وب سایت [dnslab](dnslab.ir) به منظور راحتی در دست رسی به api وب سایت.

پروژه اصلی [لینک](https://github.com/AkbarAsghari/DNSLab-WebSite)

## راهنمای نصب

برای نصب ابزار نیاز به نصب Go
دارید برای نصب به این [link](https://go.dev/doc/install) .مراجعه کنید

بعد از نصب این کامند را اجرا کنید.

> go install github.com/NoobforAl/DnsLab@latest

بعد از نصب Dnslab
چک کنید که به درستی ابزار نصب شده است.

```bash
$ dnslab

 Ipv4:  your ip
 Ipv6:
```

در صورتی که تمایلی به نصب Go
ندارید می توانید باینری فایل را برای خود دانلود کنید.
[DownloadLink](https://mega.nz/folder/AKcTWJCS#Y5FU8rIEy9ZlBRsyIfpCPA)

توجه: برای تمامی سیستم عامل ها تعبیه نشده است
و ممکن است برای برخی از معماری پردازنده ها درست کار نکند.

# راهنمای CLI

| کامند |                                         توضیحات                                          |
| ----- | :--------------------------------------------------------------------------------------: |
| -h    |                                       دیدن راهنما                                        |
| -dl   |                                  دیدن ای پی با دی ان اس                                  |
| -ip   |                  دیدن ای پی به صورت پیش فرض ای پی شما نمایش داده می شود                  |
| -op   |                                   چک کردن پورت ها باز                                    |
| -pi   |                        پینگ کردن ای پی خود یا مقداری که وارد شده                         |
| -rl   |                                  دیدن دی ان اس با ای پی                                  |
| -t    |                توکنی که از وب سایت دریافت کرده اید با این دستور وارد کنید                |
| -ts   |                             زمان مرتب چک کردن ای پی به ساعت                              |
| -uip  |                   اپدیت کردن ای پی به شرطی که توکنی را وارد کرده باشید                   |
| -up   | چک کردن مداوم ای پی بسته به زمان تنظیم شده به صورت پیش فرض بر هر 3 ساعت یک بار تنظیم است |

# Example

دیدن ip با DNS

```bash
$ dnslab -dl

 Ipv4:  your ip
 Ipv6:
Enter dns: developer.google.com
 Ip:  172.217.169.110

```

آپدیت کردن IP

```bash
$ dnslab -uip -t your token

 Ipv4:  5.53.56.41
 Ipv6:
 IP updated:  true
 IP updated:  true
```

</div>
