FROM golang:alpine AS builder
ARG TARGETOS
ARG TARGETARCH
WORKDIR $GOPATH/src/github.com/nihaals/bottom-go
COPY . .
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-w -s" -o /go/bin/bottom cmd/bottom/main.go

FROM scratch
COPY --from=builder /go/bin/bottom /
ENTRYPOINT ["/bottom"]
