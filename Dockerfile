FROM golang:latest

RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/go-bongo/bongo
RUN go get github.com/getsentry/sentry-go/gin

RUN export PATH="$PATH:$GOPATH/bin"

WORKDIR /api/
