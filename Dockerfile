FROM scratch
LABEL maintainer "Ivan Krutov <vania-pooh@aerokube.com>"

COPY selenium-mock /

EXPOSE 4444
CMD ["/selenium-mock"]
