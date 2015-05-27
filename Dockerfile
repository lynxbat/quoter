FROM golang

ADD . /go/src/github.com/lynxbat/quoter
RUN go install github.com/golang/lynxbat/quoter
ENTRYPOINT /go/bin/quoter
EXPOSE 8080