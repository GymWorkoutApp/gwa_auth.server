FROM golang:alpine

WORKDIR /go/src/github.com/GymWorkoutApp/gwa_common.server

RUN apk add --update --no-cache \
    make \
    build-base \
    jq \
    curl \
    tzdata \
    git \
    libffi-dev \
    postgresql-dev \
    gcc g++ \
    ca-certificates && \
    update-ca-certificates

ADD . /go/src/github.com/GymWorkoutApp/gwa_common.server

RUN cp /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime && \
    echo "America/Sao_Paulo" > /etc/timezone

EXPOSE 8080

RUN apk update && \
    apk add glide && \
    glide install && \
    go build