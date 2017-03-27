FROM scratch
LABEL maintainer "Ivan Krutov <vania-pooh@aerokube.com>"

COPY selenium-mock /usr/bin

EXPOSE $PORT
CMD ["/usr/bin/selenium-mock"]
