package youtubeclient

import (
	"context"
	"testing"
)

func Test_GetThumbnail(t *testing.T) {
	tr := New("")
	thumbnailURL, err := tr.GetThumbnail(context.Background(), "")
	if err != nil {
		t.Error(err)
	}

	t.Log(thumbnailURL)
}
