package grpc

import (
	"context"
	"fmt"
	"log"
	v1 "protos-service/protos/gen/payment/protobuf"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PaymentClient struct {
	client v1.PaymentServiceClient
}

func NewPaymentClient(port int) (*PaymentClient, error) {
	conn, err := grpc.NewClient("localhost:"+strconv.Itoa(port),
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
		return nil, fmt.Errorf("err connection gprc payment service: %v", err)
	}

	return &PaymentClient{
		client: v1.NewPaymentServiceClient(conn),
	}, nil
}
