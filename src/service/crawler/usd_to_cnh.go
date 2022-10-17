package crawler

import (
	"encoding/json"
	"github.com/gocolly/colly"
	"io"
	"net/http"
	"strconv"
)

func USDToCNH() (string, string, string) {
	var price, change, percent string
	collector := colly.NewCollector()
	collector.OnHTML("span[data-test='instrument-price-last']", func(element *colly.HTMLElement) {
		price = element.Text
	})
	collector.OnHTML("span[data-test='instrument-price-change']", func(element *colly.HTMLElement) {
		change = element.Text
	})
	collector.OnHTML("span[data-test='instrument-price-change-percent']", func(element *colly.HTMLElement) {
		text := element.Text
		text = text[1:][:len(text)-2]
		percent = text
	})

	err := collector.Visit("https://cn.investing.com/currencies/usd-cnh")
	if err != nil {
		println(err.Error())
	}
	return price, change, percent
}

type dataItem struct {
	Datetime string
	Close    float32
	Open     float32
	Highest  float32
	Lowest   float32
	Amount   float32
	Percent  float32
}

func USDToCNHByDaily() []dataItem {

	url := `https://api.investing.com/api/financialdata/961728/historical/chart/?period=P1D&interval=PT5M&pointscount=120`
	req, _ := http.NewRequest("get", url, nil)
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}

	resultData, _ := io.ReadAll(resp.Body)
	jsonStr := string(resultData)
	println("", jsonStr)

	resultMap := make(map[string][][]string)
	_ = json.Unmarshal(resultData, &resultMap)
	list := resultMap["data"]

	var dailyRows []dataItem
	for _, dataRow := range list {
		closePrice, _ := strconv.ParseFloat(dataRow[1], 32)
		openPrice, _ := strconv.ParseFloat(dataRow[2], 32)
		highest, _ := strconv.ParseFloat(dataRow[3], 32)
		lowest, _ := strconv.ParseFloat(dataRow[4], 32)
		amount, _ := strconv.ParseFloat(dataRow[5], 32)
		percent, _ := strconv.ParseFloat(dataRow[6], 32)

		dailyRows = append(dailyRows, dataItem{
			Datetime: dataRow[0],
			Close:    float32(closePrice),
			Open:     float32(openPrice),
			Highest:  float32(highest),
			Lowest:   float32(lowest),
			Amount:   float32(amount),
			Percent:  float32(percent),
		})
	}
	return dailyRows
}

func T() {
	collector := colly.NewCollector()
	collector.OnResponse(func(res *colly.Response) {
		println(string(res.Body))

	})

	collector.OnError(func(res *colly.Response, err error) {
		println("encounter error:", err.Error())
		println(string(res.Body))
	})

	collector.Visit(`https://api.investing.com/api/financialdata/961728/historical/chart/?period=P1D&interval=PT5M&pointscount=120`)
}
