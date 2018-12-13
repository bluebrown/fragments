FROM golang:1.11

WORKDIR /$GOPATH/src/github.com/bluebrown/fragments/

COPY ./vendor ./vendor