package main

import (
	"github.com/gocolly/colly"
)

type Work struct {
	wname, money, company string
}

func main() {
	c := colly.NewCollector()
	var works []Work
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	c.OnHTML(".magritte-redesign.vacancy-info--ieHKDTkezpEj0Gsx", func(e *colly.HTMLElement) {
		wname := e.ChildText("magritte-text___tkzIl_6-0-4")
		money := e.ChildText("magritte-text___pbpft_4-1-1 magritte-text_style-primary___AQ7MW_4-1-1 magritte-text_typography-label-1-regular___pi3R-_4-1-1")
		company := e.ChildText("magritte-text___pbpft_4-1-1 magritte-text_style-primary___AQ7MW_4-1-1 magritte-text_typography-label-1-regular___pi3R-_4-1-1.magritte-text___tkzIl_6-0-4")

		if wname != "" || money != "" || company != "" {
			work := Work{

				wname: wname,

				money: money,

				company: company,
			}

			works = append(works, work)

		}
	})
	c.Visit("https://hh.ru/search/vacancy?text=&area=1&hhtmFrom=main&hhtmFromLabel=vacancy_search_line")
}
