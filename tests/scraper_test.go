package test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/qascade/yast/scraper"
)

func TestScraper(t *testing.T) {
	fmt.Println("Executing Scraper Test")
	var tests []string = []string{
		"spiderman",
		"captain america",
		"avengers",
		"john wick",
	}
	for _, m := range tests {
		queryContext := scraper.NewQueryContext("movie", m, scraper.TARGET_1337X)

		allowedDomain := scraper.GetAllowedDomain(scraper.TARGET_1337X)
		require.Equal(t, "1337x.to", allowedDomain)

		testScraper := scraper.NewScraper(allowedDomain)

		results, err := testScraper.Scrape(queryContext)

		require.NoError(t, err, fmt.Sprintf("error scraping: %s", err))
		require.NotEmpty(t, results, "result is empty")
	}
}
