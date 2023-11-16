FROM golang:1.19-alpine as build

RUN apk update && apk add ca-certificates && apk add tzdata && apk add build-base

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .
RUN cp /build/.env .
RUN cp /build/supervisord.conf .
RUN cp -r /build/data .

FROM bash as runtime

RUN apk add curl

WORKDIR /app

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /dist/main .
COPY --from=build /build/.env .
COPY --from=build /build/supervisord.conf /etc
COPY --from=build /build/data ./data

EXPOSE 8000 9001

ENV TZ Asia/Jakarta

COPY --from=ochinchina/supervisord:latest /usr/local/bin/supervisord /usr/local/bin/supervisord

CMD ["/usr/local/bin/supervisord"]