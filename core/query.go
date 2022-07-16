package core

import(
	"github.com/qascade/yast/scraper"
)

// type Query interface {
// 	Search()
// 	//GetResults() 
// }

type SearchQuery struct{
	Searched bool //To check whether the query has been searched or not
	Results []*scraper.Result
	Context *scraper.QueryContext
}


func NewSearchQuery(context *scraper.QueryContext) *SearchQuery{
	return &SearchQuery{
		Searched: false,
		Context: context,
	}
}
func (q *SearchQuery) Search() (results []*scraper.Result, err error){
	scraper := scraper.NewScraper()
	results, err = scraper.Scrape(q.Context)
	return
}