/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/
package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
	//"github.com/qascade/yast/movie"
)

const TARGET_PIRATEBAY = "pirate-bay"
const TARGET_1337X = "1337x"
func GetAllowedDomain(target string) string {
	if target == TARGET_PIRATEBAY {
		return "proxifiedpiratebay.org"
	}
	if target == TARGET_1337X {
		return "1337x.to"
	}
	return ""
}
func (s *Scraper) scrapePirateBay(context *QueryContext) (results []Result, err error) {
	// results = make([]Result, 0)
	// url := fmt.Sprintf("https://1337x.to/search/spiderman/1/")
	// 	resp, err := http.Get(url)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	//  //We Read the response body on the line below.
	// 	body, err := ioutil.ReadAll(resp.Body)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	//  //Convert the body to type string
	// 	sb := string(body)
	// 	fmt.Println(sb)
	return nil ,nil
}

func (s *Scraper) scrape1337x(context *QueryContext) (results []Result, err error) {
	results = make([]Result, 0)
	url := fmt.Sprintf("https://1337x.to/search/%s/1/", context.Query)
	s.collector.OnHTML("tr", func(e *colly.HTMLElement){
		movieMetaDataDOM := e.DOM
		fmt.Println(movieMetaDataDOM.Text())

	})
	s.collector.Visit(url)
	return
}