FROM gocv/opencv:4.5.1 AS builder

ENV GOPATH /go

COPY . /go/src/gocv.io/x/gocv/

COPY . /go/src/image-proc/

WORKDIR /go/src/image-proc
RUN go build -o /build/image-proc ./cmd/image-proc/

CMD ["/build/image-proc"]