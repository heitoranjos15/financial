FROM golang:1.7-alpine

RUN apk add --no-cache --update \
    gcc \
    musl-dev \
    git \
    openssh \
    openssh-client \
    curl

# ADD . /go/app

WORKDIR /go/app

# RUN \
#        apk add --no-cache bash git openssh && \
       # go get -u github.com/minio/minio-go

CMD ["go","run","main.go"]
