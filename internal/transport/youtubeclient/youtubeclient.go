package youtubeclient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sletkov/thumbnail-proxy/internal/transport/youtubeclient/utils"
	"go.uber.org/zap"
	"net/http"
)

type ThumbnailTransport struct {
	youtubeAPIKey string
}

func New(youtubeAPIKey string) *ThumbnailTransport {
	return &ThumbnailTransport{
		youtubeAPIKey: youtubeAPIKey,
	}
}

type VideoResponse struct {
	Items []struct {
		Snippet struct {
			Thumbnails struct {
				Default struct {
					URL string `json:"url"`
				} `json:"default"`
				Medium struct {
					URL string `json:"url"`
				} `json:"medium"`
				High struct {
					URL string `json:"url"`
				} `json:"high"`
			} `json:"thumbnails"`
		} `json:"snippet"`
	} `json:"items"`
}

func (tr *ThumbnailTransport) GetThumbnail(ctx context.Context, URL string) (string, error) {
	videoID, err := utils.GetVideoIDFromURL(URL)
	if err != nil {
		zap.L().Error("youtube-client [GetThumbnail]", zap.Error(err))
		return "", err
	}

	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?id=%s&key=%s&part=snippet", videoID, tr.youtubeAPIKey)

	resp, err := http.Get(url)
	if err != nil {
		zap.L().Error("youtube-client [GetThumbnail]", zap.Error(err))
		return "", err
	}
	defer resp.Body.Close()

	var videoResponse VideoResponse
	if err := json.NewDecoder(resp.Body).Decode(&videoResponse); err != nil {
		zap.L().Error("youtube-client [GetThumbnail]", zap.Error(err))
		return "", err
	}

	if len(videoResponse.Items) == 0 {
		err := fmt.Errorf("video not found")
		zap.L().Error("youtube-client [GetThumbnail]", zap.Error(err))

		return "", err
	}

	return videoResponse.Items[0].Snippet.Thumbnails.High.URL, nil
}
