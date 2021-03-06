package core

import(
	"fmt"
	"github.com/qascade/yast/scraper"
)
//Keeping this here for now, might be used in building Query Registry
type Query interface {
	Search()
	GetResults() 
}

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
	if err != nil {
		return nil, err
	}
	q.Results = results
	q.Searched = true
	return	
}
func (q *SearchQuery) GetResults() (results []*scraper.Result, err error){
	if !q.Searched {
		err = fmt.Errorf("Query has not been searched yet")
		return
	}
	return q.Results, nil
}