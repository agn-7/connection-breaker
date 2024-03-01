## Connection Breaker Server

A RESTful API service built with Go and Gin framework, providing endpoints to control certain system operations such as disabling the internet and initiating system shutdown on a Windows server.

### Build

If the respective machine is linux, before building set the following enviromnent variables:

```bash
export GOARCH=amd64
export GOOS=windows
```

```bash
go build
```
