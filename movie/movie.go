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

// Title() and Description() are a part of the bubbles.list.Item interface.
//Need these methods to render list item in the listmodel view.
// Removing these functions will cause the listmodel view to not render properly.

func (m Movie) Title() string {
	return m.Name
}

//TODO: Modify this to show Metadata for the result item
func (m Movie) Description() string {
	return "This is a movie stub"
}
