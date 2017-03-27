FROM ubuntu:16.04
MAINTAINER Ivan Krutov <vania-pooh@aerokube.com>

COPY selenium-mock /usr/bin

EXPOSE $PORT
CMD ["/usr/bin/selenium-mock"]
