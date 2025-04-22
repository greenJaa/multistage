FROM golang:1.23 AS builder

WORKDIR /go/src/github.com/alexellis/href-counter/

# Copy your source code first
COPY app.go ./

# Initialize go.mod (only if not present already)
RUN go mod init go-app || true

# Download dependencies (this includes html if it's in app.go)
RUN go mod tidy

# Build the app
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o app

# Final image
FROM alpine:3.18
RUN apk update && apk add --no-cache ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/alexellis/href-counter/app .

CMD ["./app"]

