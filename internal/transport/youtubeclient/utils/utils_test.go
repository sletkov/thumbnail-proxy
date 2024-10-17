package utils

import (
	"fmt"
	"testing"
)

func Test_GetVideoIDFromURL(t *testing.T) {
	testCases := []struct {
		name    string
		url     string
		want    string
		wantErr bool
	}{
		{
			name:    "empty",
			url:     "",
			want:    "",
			wantErr: true,
		},
		{
			name:    "not youtubeclient link",
			url:     "https://www.google.com/",
			want:    "",
			wantErr: true,
		},
		{
			name:    "OK",
			url:     "https://www.youtube.com/watch?v=NbTSXG6iAO8",
			want:    "NbTSXG6iAO8",
			wantErr: false,
		},

		{
			name:    "youtubeclient link without id",
			url:     "https://www.youtube.com",
			want:    "",
			wantErr: true,
		},

		{
			name:    "not youtubeclient link with id",
			url:     "https://www.google.com?v=NbTSXG6iAO8",
			want:    "",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			id, err := GetVideoIDFromURL(tc.url)

			if err != nil && tc.wantErr == false {
				t.Error(err)
				return
			}

			if id != tc.want {
				t.Error(fmt.Sprintf("not expected id: got %s, want %s", id, tc.want))
				return
			}
		})

	}
}
