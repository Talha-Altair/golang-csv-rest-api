FROM golang:latest

ADD . /app

WORKDIR /app

RUN go get -d ./...

RUN go install

CMD ["talha"]
