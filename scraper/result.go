package scraper

// Interface that accepts any type of result coming from Scraper
type Result interface {
}

type Results []*Result
