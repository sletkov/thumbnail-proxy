package service

import (
	"context"
	"go.uber.org/zap"
)

type ThumbnailCache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
}

type ThumbnailTransport interface {
	GetThumbnail(ctx context.Context, URL string) (string, error)
}

type ThumbnailService struct {
	ThumbnailCache     ThumbnailCache
	ThumbnailTransport ThumbnailTransport
}

func New(cache ThumbnailCache, transport ThumbnailTransport) *ThumbnailService {
	return &ThumbnailService{
		ThumbnailCache:     cache,
		ThumbnailTransport: transport,
	}
}

func (s *ThumbnailService) GetThumbnail(ctx context.Context, URL []string) ([]string, error) {
	result := make([]string, 0, len(URL))
	for _, url := range URL {
		thumbnailURL, err := s.ThumbnailCache.Get(ctx, url)
		if err != nil {
			zap.L().Warn("thumbnail link not found in cache", zap.Error(err))

			urlFromTransport, err := s.ThumbnailTransport.GetThumbnail(ctx, url)
			if err != nil {
				zap.L().Error("[ThumbnailService] GetThumbnail", zap.Error(err))
				continue
			}

			result = append(result, urlFromTransport)

			err = s.ThumbnailCache.Set(ctx, url, urlFromTransport)
			if err != nil {
				zap.L().Error("[ThumbnailService] GetThumbnail", zap.Error(err))
				continue
			}

			continue
		}

		result = append(result, thumbnailURL)
	}

	return result, nil
}
