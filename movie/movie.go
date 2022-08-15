/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/
package movie

import (
	"fmt"
	"github.com/qascade/yast/utils"
)

type Movie struct {
	Name     string
	Uploaded string
	Magnet   string //link webtorrent
	Size     string
	Seeds    string //as of new seeds are already sorted in descending order in 1337x.to
	Uploader string
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

// TODO: Modify this to show Metadata for the result item
func (m Movie) Description() string {
	var metadata = fmt.Sprintf("Uploaded: %s || Size: %s || Seeds: %s", m.Uploaded, m.Size, m.Seeds)
	return metadata
}
