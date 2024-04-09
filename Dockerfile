FROM golang:1.22.2 AS build-env

ADD . /dockerdev
WORKDIR /dockerdev
EXPOSE 8000:8000

RUN go build -o /server

CMD ["/server"]
