FROM golang:1.26-alpine AS build
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o server ./cmd/api

FROM alpine:3.24
RUN apk add --no-cache ca-certificates
COPY --from=build /app/server /server
EXPOSE 8080
CMD ["/server"]