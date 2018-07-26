FROM golang:1.10.3 as builder

WORKDIR /go/src/simple-container-app
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o simple-container-app .

FROM alpine
WORKDIR /root/

COPY --from=builder /go/src/simple-container-app/simple-container-app .

CMD ["./simple-container-app"]

