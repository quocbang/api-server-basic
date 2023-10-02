FROM golang:1.20-alpine3.18 as server

RUN apk update && apk upgrade
RUN apk add --no-cache git
RUN apk add --update --no-cache curl build-base
RUN apk update && apk add openssh

ARG TEST_DB_ADDRESS
ARG TEST_DB_DATBASE
ARG TEST_DB_SCHEMA
ARG TEST_DB_USERNAME
ARG TEST_DB_PASSWORD
ARG TEST_REDIS_ADDRESS
ARG TEST_REDIS_PASSWORD
ARG TEST_REDIS_DATABASE
ARG TEST_SECRET_KEY

ENV SERVER_DIR ${GOPATH}/src/github.com/quocbang/api-server-basic

COPY ./ ${SERVER_DIR}

WORKDIR ${SERVER_DIR}

# set data connection for test units.
ENV DB_ADDRESS=DB_TEST_ADDRESS
ENV DB_PORT=5432
ENV DB_DATABASE=DB_TEST_NAME
ENV DB_SCHEMA=DB_TEST_SCHEMA
ENV DB_USERNAME=DB_TEST_USERNAME
ENV DB_PASSWORD=DB_TEST_PASSWORD

ENV REDIS_ADDRESS=TEST_REDIS_ADDRESS
ENV REDIS_PASSWORD=TEST_REDIS_PASSWORD
ENV REDIS_DATABASE=TEST_REDIS_DATABASE

ENV SECRET_KEY=TEST_SECRET_KEY

RUN go generate .
RUN go mod download
RUN go vet ./...
# RUN go test ./...
RUN go build -race -ldflags "-extldflags '-static'" -o /opt/api-server-basic/server .

CMD ["/bin/sh"]

FROM alpine:latest

WORKDIR /root/

COPY --from=server /opt/api-server-basic/server /root/server
