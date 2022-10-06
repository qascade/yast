package scraper

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/qascade/yast/movie"

	colly "github.com/gocolly/colly"
)

// scrape1337x returns the scraped results in Result type
func (s *Scraper) scrape1337x(context *QueryContext) ([]Result, error) {

	var links []string
	var movieList []movie.Movie
	var results []Result

	url := fmt.Sprintf("https://1337x.to/search/%s/1/", context.Query)

	s.collector.OnHTML("tr", func(e *colly.HTMLElement) {

		doc := e.DOM
		anchorTags := doc.Find("a").Siblings()

		if attr, exists := anchorTags.Attr("href"); exists {
			if strings.Contains(attr, "/torrent/") && !strings.HasPrefix(attr, "http://") {
				link := "https://1337x.to" + attr
				links = append(links, link)
			}
		}
		newmovie := movieFromString(doc)
		movieList = append(movieList, newmovie)

	})
	s.collector.OnHTML("div.col-9 div.box-info", func(h *colly.HTMLElement) {

		metadata := h.DOM

		movieName := metadata.Find("div.box-info-heading h1").Text()
		metadata.Find("a").Each(func(i int, s *goquery.Selection) {
			if magnet, exists := s.Attr("href"); exists {
				if strings.Contains(magnet, "magnet") {
					for i, movie := range movieList {
						if compareNamesFromDiffPages(strings.TrimSpace(movie.Name), strings.TrimSpace(movieName)) {
							movieList[i].Magnet = magnet
						}
					}
				}
			}
		})
	})

	s.collector.Visit(url)
	for _, link := range links {
		s.collector.Visit(link)
	}
	for _, movie := range movieList[1:] {
		results = append(results, movie)
	}
	return results, nil
}

// movieFromString returns a type Movie
func movieFromString(document *goquery.Selection) movie.Movie {
	return movie.Movie{
		Name:     document.Find("td.name > a:nth-child(2)").Text(),
		Uploaded: document.Find("td.coll-date").Text(),
		Size:     formatSize1337x(strings.Replace(document.Find("td.size").Text(), document.Find("td.seeds").Text(), "", -1)),
		Seeds:    document.Find("td.seeds").Text(),
		Uploader: document.Find("td.coll-5").Text(),
	}
}

// compareNamesFromDiffPages compares the movie name on different pages
func compareNamesFromDiffPages(a string, b string) bool {

	if len(a) == 0 || len(b) == 0 {
		return false
	}
	limit := len(a)
	if len(b) < len(a) {
		limit = len(b)
	}
	if limit > 65 {
		limit = 65
	}
	for i := 0; i < limit; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func formatSize1337x(size string) string {
	idx := strings.Index(size, "B")
	size = size[:idx+1]
	return size
}
