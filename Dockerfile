FROM golang:alpine AS builder
WORKDIR $GOPATH/src/github.com/nihaals/bottom-go
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/bottom cmd/main.go

FROM scratch
COPY --from=builder /go/bin/bottom /
ENTRYPOINT ["/bottom"]
