FROM golang:1.10 as builder
WORKDIR /go/src/microservice
RUN go get -u github.com/golang/dep/cmd/dep
COPY .  .
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/microservice .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/bin/microservice .
COPY cert/ cert/
EXPOSE 80
CMD ["./microservice"]