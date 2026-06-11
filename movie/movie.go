/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/
package movie

import "fmt"

type Movie struct {
	Name     string
	Uploaded string
	Magnet   string
	Size     string
	Seeds    string
	Uploader string
	Quality  string
}

func (m Movie) FilterValue() string {
	return fmt.Sprintf("%s %s %s %s", m.Name, m.Quality, m.Size, m.Seeds)
}

func (m Movie) Title() string {
	return m.Name
}

func (m Movie) Description() string {
	return fmt.Sprintf("Quality: %s || Uploaded: %s || Size: %s || Seeds: %s", m.Quality, m.Uploaded, m.Size, m.Seeds)
}
