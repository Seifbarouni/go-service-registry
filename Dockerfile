FROM golang:alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download

RUN go build -o main .

EXPOSE 8761

CMD ["/app/main"]

