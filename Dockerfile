FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY google-oauth google-oauth
COPY config.production.tml config.tml

EXPOSE 8080

CMD ./google-oauth
