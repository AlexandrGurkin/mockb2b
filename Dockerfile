FROM golang:1.12.5-alpine3.9 as builder
RUN mkdir /smsServer && \
  apk add --no-cache git
WORKDIR /smsServer
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ./cmd/smsServer/

FROM scratch
COPY --from=builder /smsServer/smsServer /smsServer
COPY /config/config.json /etc/sms_server/config.json

EXPOSE 8082
CMD ["/smsServer"]