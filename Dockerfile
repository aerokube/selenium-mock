FROM ubuntu:16.04
MAINTAINER Ivan Krutov <vania-pooh@vania-pooh.com>

ENV PORT 4444

COPY selenium-mock /usr/bin
COPY entrypoint.sh /

EXPOSE 4444
ENTRYPOINT /entrypoint.sh
