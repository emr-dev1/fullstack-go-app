FROM golang:alpine as builder

# ENV GO111MODULE=on

# Installing git for fetching dependencies
RUN apk update && apk add --no-cache git

WORKDIR /app

# Copy the go mod and sum files
COPY go.mod go.sum ./

# Download deps
RUN go mod download

# Copy source from current directory to working directory inside Container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy pre-built binary file from previous stage.
# Also copy .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./main"]
