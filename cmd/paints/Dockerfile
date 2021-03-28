FROM gocv/opencv:4.5.1 AS builder

ENV GOPATH /go

COPY . /go/src/gocv.io/x/gocv/

COPY . /go/src/paints/

WORKDIR /go/src/paints
RUN go build -o /build/paints ./cmd/paints/

HEALTHCHECK --interval=5s --timeout=10s --retries=3 CMD curl -sS http://$HOSTNAME:8000/health-check || exit 1

CMD ["/build/paints"]