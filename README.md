# Selenium Mock
This repository contains a small binary emulating real Selenium Hub responses. The main application - load testing of Selenium proxies like [Gridrouter](http://github.com/aandryashin/ggr).

## Building & Running
To build install [Govendor](https://github.com/kardianos/govendor) and type:
```
$ govendor sync
$ go build
```
Then run:
```
$ ./mock -port 4444
```
To create Docker container type:
```
$ docker build -t selenium-mock:latest .
```
