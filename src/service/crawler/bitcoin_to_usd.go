package crawler

import "github.com/gocolly/colly"

func BitcoinToCNH() (string, string, string) {
	var price, change, percent string
	collector := colly.NewCollector()
	collector.OnHTML("span[class='pid-1057391-last'][id='last_last']", func(element *colly.HTMLElement) {
		price = element.Text
	})
	collector.OnHTML("span[class~='pid-1057391-pc']", func(element *colly.HTMLElement) {
		change = element.Text
	})
	collector.OnHTML("span[class~='pid-1057391-pcp']", func(element *colly.HTMLElement) {
		percent = element.Text
	})
	_ = collector.Visit(`https://cn.investing.com/crypto/bitcoin`)
	return price, change, percent
}
