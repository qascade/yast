package test

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/require"

	"github.com/qascade/yast/scraper"
)

func TestScraper(t *testing.T) {
	fmt.Println("Executing Scraper Test")
	queryContext := scraper.NewQueryContext("movie", "spiderman", scraper.TARGET_1337X)

	allowedDomain := scraper.GetAllowedDomain(scraper.TARGET_1337X)
	require.Equal(t, "1337x.to", allowedDomain)

	testScraper := scraper.NewScraper(allowedDomain)

	results, err := testScraper.Scrape(queryContext)

	fmt.Printf("%+v\n", results)

	require.NoError(t, err, fmt.Sprintf("error scraping: %s", err))
	require.NotEmpty(t, results, fmt.Sprintf("result is empty"))
}