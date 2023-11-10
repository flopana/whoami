FROM golang:1.21

COPY ./ ./

RUN go version
RUN go get
RUN go build -o /bin/whoami main.go

FROM gcr.io/distroless/static-debian12

COPY --from=0 /bin/whoami /bin/whoami
CMD ["/bin/hello"]
