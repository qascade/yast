/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/
package scraper

import (
	colly "github.com/gocolly/colly"
	"github.com/qascade/yast/utils"
)

type Scraper struct {
	collector *colly.Collector
}

type QueryContext struct {
	Type   string //movie, series
	Query  string
	Target string
}

func NewScraper(domain string) *Scraper {
	collector := colly.NewCollector(
		colly.AllowedDomains(domain),
	)
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
	//results = make([]Result, 0)
	if context.Type == "movie" {
		if context.Target == TARGET_PIRATEBAY {
			results, err = s.scrapePirateBay(context)
			return 
		}
		if context.Target == TARGET_1337X {
			//utils.LogTraceMsg("1337x scraper not yet implemented")
			results, err = s.scrape1337x(context)
			//results, err = s.scrape1337x(context)
			return 
		}
	} 
	if context.Type == "series" {
		utils.LogTraceMsg("series scraper not implemented yet")
	}

	return
}
