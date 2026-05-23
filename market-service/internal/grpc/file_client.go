package grpc

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type FileClient struct {
}

const serviceConfig = `{
	"loadBalancingConfig": [{ "round_robin": {} }],
	"methodConfig": [{
		"name": [{
			"method": "Get",
			"service": "news.v1.NewsService"
		}],
		"retryPolicy": {
			"backoffMultiplier": 1.5,
			"initialBackoff": "0.1s",
			"maxAttempts": 5,
			"maxBackoff": "0.5s",
			"retryableStatusCodes": ["INTERNAL","UNAVAILABLE"]
		},
		"timeout": "2s",
		"waitForReady": true
	}]
}`

func NewFileClient(url string) (*FileClient, error) {
	conn, err := grpc.NewClient(url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
				log.Println("unary interceptor called")
				return invoker(ctx, method, req, reply, cc, opts...)
			},
		),
		grpc.WithChainStreamInterceptor(
			func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) { //nolint:lll // Refactor to create a separate interceptor.
				log.Println("stream interceptor called")
				return streamer(ctx, desc, cc, method, opts...)
			},
		),

		grpc.WithDefaultServiceConfig(serviceConfig),
	)

	if err != nil {
		return nil, fmt.Errorf("err connection gprc file service: %v", err)
	}
	client := v1.NewPaymentServiceClient()
	return &FileClient{}
}
