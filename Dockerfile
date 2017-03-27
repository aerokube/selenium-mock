FROM scratch
LABEL maintainer "Ivan Krutov <vania-pooh@aerokube.com>"

COPY selenium-mock /usr/bin

EXPOSE 4444
CMD ["/usr/bin/selenium-mock"]
