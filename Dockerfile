FROM golang:1.17.7-buster as builder

WORKDIR /app

ADD . .

RUN go mod download

RUN go build -o app-exec

FROM alpine:3.15.0 as runner

COPY --from=builder /app/app-exec /app

CMD [ "/app" ]

