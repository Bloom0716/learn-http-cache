FROM golang:1.24-alpine
RUN apk add --no-cache ca-certificates tzdata git
WORKDIR /root/app
RUN go install github.com/air-verse/air@latest
COPY ./app/go.mod ./app/go.sum ./
RUN go mod download
COPY ./app .
CMD ["air"]
