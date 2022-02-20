FROM golang:1.17.7-alpine3.15 as builder

WORKDIR /code

ADD . .

RUN go mod download

RUN go build -o app

FROM alpine:3.15.0 as runner

WORKDIR /

COPY --from=builder /code/app /usr/local/bin/app

CMD [ "app" ]