# syntax=docker/dockerfile:1
FROM golang:1.16-alpine
ENV GOOS "linux"
ENV GOARCH "arm64"
WORKDIR /app

COPY *.go ./
COPY go.mod ./
RUN go get -d -v
RUN go build -o ./go-login

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app
EXPOSE 8080
COPY --from=0 /app/go-login ./
CMD [ "./go-login" ]


