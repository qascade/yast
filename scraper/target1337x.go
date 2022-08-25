package scraper

import (
	"fmt"
	"strings"

	"github.com/qascade/yast/movie"

	colly "github.com/gocolly/colly"
)

func (s *Scraper) scrape1337x(context *QueryContext) (results []Result, err error) {

	results = make([]Result, 0)
	url := fmt.Sprintf("https://1337x.to/search/%s/1/", context.Query)
	s.collector.OnHTML("tr", func(e *colly.HTMLElement) {
		movieMetaDataDOM := e.DOM
		anchorTagDom := movieMetaDataDOM.Find("a").Siblings()

		var magnet string
		if magnetPageUrl, exists := anchorTagDom.Attr("href"); exists {
			magnet = getMagnetLink("https://1337x.to" + magnetPageUrl)
		}

		newMovie := movieFromString(movieMetaDataDOM.Text(), magnet)
		newMovie.Magnet = magnet
		results = append(results, newMovie)

	})

	s.collector.Visit(url)
	results = results[1:]
	return
}

func movieFromString(data string, magnet string) movie.Movie {
	elements := strings.Split(data, "\n")
	// Need to find a better way to do the parsing. This is fragile.
	return movie.Movie{
		Name:     elements[1],
		Uploaded: elements[4],
		Magnet:   magnet,
		Size:     formatSize1337x(elements[5]),
		Seeds:    elements[2],
		Uploader: elements[6],
	}
}

func getMagnetLink(url string) (magnet string) {
	s := NewScraper("1337x.to")

	s.collector.OnHTML("a", func(e *colly.HTMLElement) {
		dom := e.DOM
		potMagnet, _ := dom.Attr("href")
		if strings.Contains(potMagnet, "magnet") {
			magnet = potMagnet
		}
	})

	s.collector.Visit(url)
	return
}
func formatSize1337x(size string) string {
	idx := strings.Index(size, "B")
	size = size[:idx+1]
	return size
}
