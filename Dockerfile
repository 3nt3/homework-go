FROM golang

ADD . /go/src/github.com/3nt3/homework

WORKDIR /go/src/github.com/3nt3/homework

RUN go build .

ENTRYPOINT /go/src/github.com/3nt3/homework/homework

EXPOSE 8000
