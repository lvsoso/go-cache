```shell
(base) lv@lv:client$ curl 127.0.0.1:12345/cache/a -XPUT -daa
(base) lv@lv:client$ curl 127.0.0.1:12345/cache/a
aa

(base) lv@lv:client$ curl 127.0.0.1:12345/status
{"Count":1,"KeySize":1,"ValueSize":2}

# waitting...  30 s

(base) lv@lv:client$ curl 127.0.0.1:12345/status
{"Count":0,"KeySize":0,"ValueSize":0}

(base) lv@lv:client$ curl 127.0.0.1:12345/cache/a -v
*   Trying 127.0.0.1:12345...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 12345 (#0)
> GET /cache/a HTTP/1.1
> Host: 127.0.0.1:12345
> User-Agent: curl/7.65.3
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 404 Not Found
< Date: Sun, 03 Apr 2022 19:17:11 GMT
< Content-Length: 0
< 
* Connection #0 to host 127.0.0.1 left intact

```