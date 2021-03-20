package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	api "paint/pkg/api/proto_files"
)

func ExampleDrawDefaultContours(ctx context.Context, client api.ImageProcessingServiceClient) {
	req := api.ContoursRequest{}
	res, err := client.DrawDefaultContours(ctx, &req)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(res.OutPicture)
}

func main() {
	conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = conn.Close()
	}()

	client := api.NewImageProcessingServiceClient(conn)
	ExampleDrawDefaultContours(context.Background(), client)
}
