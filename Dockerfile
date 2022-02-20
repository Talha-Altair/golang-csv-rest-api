FROM golang:1.17.7-alpine3.15 as builder

WORKDIR /code

ADD . .

RUN go mod download

CMD [ "go", "run", "." ]