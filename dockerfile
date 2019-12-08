FROM golang:alpine as builder

ADD . /go/src/github.com/marine
WORKDIR /go/src/github.com/marine

RUN apk add git
RUN apk --no-cache add ca-certificates
# Building the Go executable for linux
RUN GO111MODULE=on GOOS=linux GOARCH=386 go build -o /app -i main.go

########################################################

FROM scratch

COPY firebase.json ./
COPY --from=builder /app ./
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /

ENTRYPOINT ["./app"]