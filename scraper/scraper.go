/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved. 
This Project is under BSD-3 License Clause. 
Look at License for more detail. 
*/
package scraper

type Scraper struct {
	//collector colly.NewCollector()
}

type QueryContext struct {
	Type  string //movie, series
	Query string
}

func NewScraper() *Scraper {
	return &Scraper{}
}

func NewQueryContext(flagtype string, query string) *QueryContext {
	return &QueryContext{
		Type:  flagtype,
		Query: query,
	}
}
func (s *Scraper) Scrape(context *QueryContext) (results []*Result, err error) {
	results = make([]*Result, 0)
	err = nil
	return
}
