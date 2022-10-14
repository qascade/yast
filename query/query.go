/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/
package core

import (
	"github.com/pkg/errors"
	"github.com/qascade/yast/scraper"
)

// Keeping this here for now, might be used in building Query Registry
type Query interface {
	Search() (results []scraper.Result, err error)
	//GetResults()
}

type SearchQuery struct {
	searched bool //To check whether the query has been searched or not
	Results  []scraper.Result
	Context  *scraper.QueryContext
}

func NewSearchQuery(context *scraper.QueryContext) *SearchQuery {
	return &SearchQuery{
		searched: false,
		Context:  context,
	}
}
func (q *SearchQuery) Search() (results []scraper.Result, err error) {
	allowedDomain := scraper.GetAllowedDomain(q.Context.Target)
	scraper := scraper.NewScraper(allowedDomain)
	results, err = scraper.Scrape(q.Context)
	if err != nil {
		return nil, err
	}
	q.Results = results
	q.searched = true
	return
}

// To be used when implementing Query History
func (q *SearchQuery) GetResults() (results []scraper.Result, err error) {
	if !q.searched {
		err = errors.Errorf("Query has not been searched yet")
		return
	}
	return q.Results, nil
}
