FROM golang:1.22-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o serviceServer ./cmd

RUN chmod +x /app/serviceServer


FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/serviceServer /app

CMD [ "/app/serviceServer" ]