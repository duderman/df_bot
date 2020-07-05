FROM golang:alpine as builder
WORKDIR /build
ADD go* /build/
RUN go mod download
ADD main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM scratch
WORKDIR /app
COPY --from=builder /build/main /app/
ADD ca-certificates.crt /etc/ssl/certs/
CMD ["./main"]
