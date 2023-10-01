FROM golang:1.19.2 AS development
WORKDIR /templateserver
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/cespare/reflex@latest
EXPOSE 4000
# CMD reflex -g '*.go' go run main.go --start-service
ENTRYPOINT ["reflex", "-c", "reflex.conf"]

FROM golang:1.19.2 AS builder
ENV GOOS linux
ENV CGO_ENABLED 0
WORKDIR /templateserver
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o templateserver

FROM alpine:latest AS production
RUN apk add --no-cache ca-certificates
COPY --from=builder templateserver .
EXPOSE 4000
CMD ./templateserver



