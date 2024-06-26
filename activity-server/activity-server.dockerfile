FROM golang:1.22-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o activtiyServer ./cmd

RUN chmod +x /app/activtiyServer


FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/activtiyServer /app

CMD [ "/app/activtiyServer" ]