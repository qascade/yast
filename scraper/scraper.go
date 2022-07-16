package scraper



type Scraper struct{
	//collector colly.NewCollector()
}

type QueryContext struct{
	Type string //movie, series
	Query string 
}
func NewScraper() (*Scraper){
	return &Scraper{}
}

func (s *Scraper) Scrape(context *QueryContext) (results []*Result, err error){
	return nil, nil
}
