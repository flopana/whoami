FROM golang:1.21 as build
WORKDIR /whoami

COPY ./ ./

RUN go version
RUN go get

# Strip debug information and compile only the binary
# https://golang.org/cmd/link/
# https://github.com/xaionaro/documentation/blob/master/golang/reduce-binary-size.md#1-strip
RUN go build -ldflags="-w -s" -o /go/bin/whoami main.go

FROM gcr.io/distroless/static-debian12

COPY --from=build /go/bin/whoami /
CMD ["/whoami"]
