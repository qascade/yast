package movie

import (
	"time"

	"github.com/qascade/yast/utils"
)

type Movie struct {
	Name     string
	Uploaded time.Time
	Magnet   string //link webtorrent
	Size     string
	Seeds    int
	Uploader string
}

func NewMovie(name string, uploaded time.Time, magnet string, size string, seeds int, uploader string) (*Movie, error) {
	return &Movie{
		Name:     name,
		Uploaded: uploaded,
		Magnet:   magnet,
		Size:     size,
		Seeds:    seeds,
		Uploader: uploader,
	}, nil
}

func (m Movie) FilterValue() string {
	utils.LogUnimplementedFunc()
	return ""
}
