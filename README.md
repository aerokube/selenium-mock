# Selenium Mock
[![Build Status](https://travis-ci.org/aerokube/selenium-mock.svg?branch=master)](https://travis-ci.org/aerokube/selenium-mock)
[![Release](https://img.shields.io/github/release/aerokube/selenium-mock.svg)](https://github.com/aerokube/selenium-mock/releases/latest)

This repository contains a small binary emulating real Selenium Hub responses. The main application - load testing of Selenium proxies like [Gridrouter](http://github.com/aandryashin/ggr).

## Running
```
# docker run --rm -it --name selenium-mock aerokube/selenium-mock:1.0.0
```

## Building
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
