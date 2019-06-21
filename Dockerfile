
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ./cmd/smsServer/

FROM scratch
COPY /mockb2b /mockb2b

EXPOSE 8080
CMD ["/mockb2b"]