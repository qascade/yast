package test

import (
	//"fmt"
	//"github.com/stretchr/testify/require"
	"testing"
	//"github.com/qascade/yast/scraper"
)

// This test will fail on workflows due to 1337x.to not being reachable.
// Please make sure that these tests pass locally.
func TestScraper(t *testing.T) {
	// fmt.Println("Executing Scraper Test")
	// var tests []string = []string{
	// 	"spiderman",
	// 	"captain america",
	// 	"avengers",
	// 	"john wick",
	// }
	// for _, m := range tests {
	// 	queryContext := scraper.NewQueryContext("movie", m, scraper.TARGET_1337X)

	// 	allowedDomain := scraper.GetAllowedDomain(scraper.TARGET_1337X)
	// 	require.Equal(t, "1337x.to", allowedDomain)

	// 	testScraper := scraper.NewScraper(allowedDomain)

	// 	results, err := testScraper.Scrape(queryContext)

	// 	require.NoError(t, err, fmt.Sprintf("error scraping: %s", err))
	// 	require.NotEmpty(t, results, "result is empty")
	// }
}
