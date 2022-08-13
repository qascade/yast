/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/
package scraper

import (
	colly "github.com/gocolly/colly"
)

type Scraper struct {
	collector *colly.Collector
}

type QueryContext struct {
	Type   string //movie, series
	Query  string
	Target string
}

func NewScraper() *Scraper {
	collector := colly.NewCollector()
	return &Scraper{
		collector: collector,
	}
}

func NewQueryContext(flagtype string, query string, target string) *QueryContext {
	return &QueryContext{
		Type:   flagtype,
		Query:  query,
		Target: target,
	}
}
func (s *Scraper) Scrape(context *QueryContext) (results []Result, err error) {
	results = make([]Result, 0)
	err = nil
	return
}
