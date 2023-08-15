# Use a multi-stage build for a smaller final image
FROM golang:1.19 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o notification .

# Use a minimal Alpine base image
FROM alpine:latest

COPY --from=build /app/notification /notification

CMD ["/myapp"]