package grpc

import (
	"context"
	"fmt"
	"io"
	"log"
	v1 "protos-service/protos/gen/file/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type FileClient struct {
	client v1.FileServiceClient
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

func NewFileClient(addr string) (*FileClient, error) {
	conn, err := grpc.NewClient(addr,
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

	return &FileClient{
		client: v1.NewFileServiceClient(conn),
	}, nil
}

func (c *FileClient) UploadFile(ctx context.Context, name, contentType string, content []byte) (*v1.UploadFileResponse, error) {
	req := &v1.UploadFileRequest{
		Name:        name,
		ContentType: contentType,
		Content:     content,
	}

	resp, err := c.client.UploadFile(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("upload file error: %w", err)
	}

	if !resp.Success {
		return nil, fmt.Errorf("upload failed: %s", resp.Error)
	}
	return resp, nil
}

func (c *FileClient) GetFile(ctx context.Context, fileID string) (*v1.GetFileResponse, error) {
	req := &v1.GetFileRequest{
		FileId: fileID,
	}

	stream, err := c.client.GetFile(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("get file stream error: %w", err)
	}

	var result *v1.GetFileResponse

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("stream recv error: %w", err)
		}

		if result == nil {
			result = &v1.GetFileResponse{
				FileId:      chunk.FileId,
				Name:        chunk.Name,
				ContentType: chunk.ContentType,
			}
		}

		result.Content = append(result.Content, chunk.Content...)
	}

	if result == nil {
		return nil, fmt.Errorf("empty response for file: %s", fileID)
	}
	return result, nil
}
