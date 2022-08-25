/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/
package scraper

import (
	"github.com/qascade/yast/utils"
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
	utils.LogUnimplementedFunc()
	return nil, nil
}
