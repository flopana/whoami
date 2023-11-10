FROM golang:1.21
WORKDIR /whoami

COPY ./ ./

RUN go version
RUN go get
# add flags -w -s for
RUN go build -ldflags="-w -s" -o /bin/whoami main.go

FROM gcr.io/distroless/static-debian12

COPY --from=0 /bin/whoami /bin/whoami
CMD ["/bin/hello"]
