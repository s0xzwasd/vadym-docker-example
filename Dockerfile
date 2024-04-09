FROM golang:1.16.3-buster AS build-env

ADD . /dockerdev
WORKDIR /dockerdev

RUN go build -o /server

FROM debian:buster

EXPOSE 8000

WORKDIR /
COPY --from=build-env /server /

CMD ["/server"]
