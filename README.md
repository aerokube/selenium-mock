# Selenium Mock
[![Build Status](https://travis-ci.org/aerokube/selenium-mock.svg?branch=master)](https://travis-ci.org/aerokube/selenium-mock)
[![Release](https://img.shields.io/github/release/aerokube/selenium-mock.svg)](https://github.com/aerokube/selenium-mock/releases/latest)

This repository contains a small binary emulating real Selenium Hub responses. The main application - load testing of Selenium proxies like [Gridrouter](http://github.com/aandryashin/ggr).

## Running
Run the following command to obtain running selenium-mock on port 4444:
```
# docker run --rm -it --name selenium-mock -p 4444:4444 aerokube/selenium-mock:1.0.0
```

## Building
To build install [Govendor](https://github.com/kardianos/govendor) and type:
```
$ go get github.com/aerokube/selenium-mock
$ cd $GOPATH/src/github.com/aerokube/selenium-mock
$ govendor sync
$ go build
```
Then run:
```
$ ./selenium-mock -port 4444
```
To create Docker container type:
```
$ docker build -t selenium-mock:latest .
```
