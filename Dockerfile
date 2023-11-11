FROM golang:1.21
WORKDIR /whoami

COPY ./ ./

RUN go version
RUN go get

# Strip debug information and compile only the binary
# https://golang.org/cmd/link/
# https://github.com/xaionaro/documentation/blob/master/golang/reduce-binary-size.md#1-strip
RUN go build -ldflags="-w -s" -o /bin/whoami main.go

FROM gcr.io/distroless/static-debian12

COPY --from=0 /bin/whoami whoami
CMD ["whoami"]
