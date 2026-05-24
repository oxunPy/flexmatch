package grpc

import (
	"context"
	"file-service/internal/services"
	v1 "protos-service/protos/gen/file/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FileServiceApi struct {
	service *services.FileService
	v1.UnimplementedFileServiceServer
}

func RegisterFileServiceApi(s *GrpcServer, svc *services.FileService) {
	v1.RegisterFileServiceServer(s.server, &FileServiceApi{
		service: svc,
	})
}

func (s *FileServiceApi) UploadFile(ctx context.Context, req *v1.UploadFileRequest) (*v1.UploadFileResponse, error) {
	if req.Name == "" || req.ContentType == "" || len(req.Content) == 0 {
		return nil, status.Error(codes.InvalidArgument, "name, content-type and content are required")
	}

	f, err := s.service.Upload(ctx, req.Name, req.ContentType, 1, req.Content)
	if err != nil {
		return &v1.UploadFileResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &v1.UploadFileResponse{
		FileId:  f.ID,
		Success: true,
		Url:     f.URL,
	}, nil
}

func (s *FileServiceApi) GetFile(req *v1.GetFileRequest, stream grpc.ServerStreamingServer[v1.GetFileResponse]) error {
	if req.FileId == "" {
		return status.Error(codes.InvalidArgument, "file id is required")
	}

	file, content, err := s.service.GetContent(stream.Context(), req.FileId)
	if err != nil {
		return status.Errorf(codes.NotFound, "file not found: %v", err)
	}

	mb := 1 << 20
	for i := 0; i < len(content); i += mb {
		end := i + mb
		if end > len(content) {
			end = len(content)
		}

		if err := stream.Send(&v1.GetFileResponse{
			FileId:      file.ID,
			Content:     content[i:end],
			ContentType: file.ContentType,
			Name:        file.Name,
		}); err != nil {
			return status.Errorf(codes.Internal, "stream sending content err: %v", err)
		}
	}

	return nil
}
