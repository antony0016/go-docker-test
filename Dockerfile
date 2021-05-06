FROM golang:latest

WORKDIR /go-test
COPY . /go-test
RUN go build .

EXPOSE 3000
ENTRYPOINT ["/go-test/go-test"]
