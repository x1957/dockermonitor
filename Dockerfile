FROM hub.c.163.com/library/golang:alpine AS builder

ADD glide.yaml /
RUN apk update && \
    apk add git && \
    go get github.com/Masterminds/glide
ADD . /go/src/github.com/x1957/dockermonitor

RUN cd /go/src/github.com/x1957/dockermonitor && \
    glide install && \
    go install github.com/x1957/dockermonitor

CMD /go/bin/dockermonitor
