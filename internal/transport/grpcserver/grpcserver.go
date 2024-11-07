package grpcserver

import (
	"context"
	proto "github.com/sletkov/thumbnail-proxy/pkg/sdk/go/thumbnailproxy_grpc"
	"go.uber.org/zap"
)

type ThumbnailProxyService interface {
	GetThumbnail(ctx context.Context, URL []string) ([]string, error)
}

type Server struct {
	proto.UnimplementedThumbnailProxyServer
	ThumbnailProxyService ThumbnailProxyService
}

func New(thumbnailProxyService ThumbnailProxyService) *Server {
	return &Server{
		ThumbnailProxyService: thumbnailProxyService,
	}
}

func (s *Server) GetThumbnail(ctx context.Context, req *proto.URLRequest) (*proto.URLResponse, error) {
	zap.L().Info("GetThumbnail was invoked")

	url, err := s.ThumbnailProxyService.GetThumbnail(ctx, req.URL)
	if err != nil {
		zap.L().Error("grpc-client [GetThumbnail]", zap.Error(err))
		return nil, err
	}

	return &proto.URLResponse{
		URL: url,
	}, nil
}
