FROM alpine:latest

WORKDIR /app

ADD genshin .

ENTRYPOINT ["/app/genshin"]