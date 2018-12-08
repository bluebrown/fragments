FROM golang:1.11.1

WORKDIR /$GOPATH/src/github.com/bluebrown/fragments/

COPY ./vendor ./vendor