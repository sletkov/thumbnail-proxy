package utils

import (
	"fmt"
	"net/url"
)

func GetVideoIDFromURL(youtubeURL string) (string, error) {
	parsedURL, err := url.Parse(youtubeURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL, %v", err)
	}

	queryParams := parsedURL.Query()
	videoID := queryParams.Get("v")

	if videoID == "" {
		return "", fmt.Errorf("failed to get id from url: %s", youtubeURL)
	}

	return videoID, nil

}
