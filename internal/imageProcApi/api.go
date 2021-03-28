package imageProcApi

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	api "paint/api/proto/pb"
	"paint/internal/cache"
	"paint/internal/cache/memory"
	"time"
)

const (
	cachedTime  = 5 * time.Minute
	clearedTime = 30 * time.Minute
	address     = "localhost:10000"
)

type Api interface {
	PyrMeanShiftFilter(ctx context.Context, in *api.PyrRequest) (*api.DefaultReply, error)
}

type imageProcApi struct {
	client                  *http.Client
	storage                 cache.Storage
	processingServiceClient api.ImageProcessingServiceClient
}

func newImageProcessingClient() api.ImageProcessingServiceClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return api.NewImageProcessingServiceClient(conn)
}

func New() Api {
	defaultTransport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:        20,
		MaxIdleConnsPerHost: 20,
		TLSHandshakeTimeout: 15 * time.Second,
	}

	client := &http.Client{
		Transport: defaultTransport,
		Timeout:   15 * time.Second,
	}

	return &imageProcApi{
		client:                  client,
		storage:                 memory.InitCash(cachedTime, clearedTime),
		processingServiceClient: newImageProcessingClient(),
	}
}
