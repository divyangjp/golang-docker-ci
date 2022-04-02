FROM golang:1.17 as base

WORKDIR /go/src/app
COPY *.go .
COPY go.* .

RUN go mod download
RUN go get -d -v ./...
RUN go vet -v
RUN go test -v

RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /go/bin/app

FROM gcr.io/distroless/static

COPY --from=base /go/bin/app /

EXPOSE 8080
CMD ["/app"]
