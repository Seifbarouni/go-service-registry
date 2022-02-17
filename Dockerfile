FROM golang:alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download

RUN go build -o service_registry .

EXPOSE 8761

CMD ["/app/service_registry"]

