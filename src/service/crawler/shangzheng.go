package crawler

import (
	"encoding/json"
	"github.com/gocolly/colly"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func ShangzhengByInverst(price *string, change *string, percent *string) {
	collector := colly.NewCollector()

	collector.OnHTML("span[data-test='instrument-price-last']", func(element *colly.HTMLElement) {
		*price = element.Text
	})

	collector.OnHTML("span[data-test='instrument-price-change']", func(element *colly.HTMLElement) {
		*change = element.Text
	})
	collector.OnHTML("span[data-test='instrument-price-change-percent']", func(element *colly.HTMLElement) {
		text := element.Text
		*percent = text[1:][:len(text)-2]
	})
	_ = collector.Visit("https://cn.investing.com/indices/shanghai-composite")
}

func Shangzheng() (string, string, string) {
	url := `http://53.push2his.eastmoney.com/api/qt/stock/kline/get?cb=jQuery3510016729102546818142_1665632758058&secid=1.000001&ut=fa5fd1943c7b386f172d6893dbfba10b&fields1=f1%2Cf2%2Cf3%2Cf4%2Cf5%2Cf6&fields2=f51%2Cf52%2Cf53%2Cf54%2Cf55%2Cf56%2Cf57%2Cf58%2Cf59%2Cf60%2Cf61&klt=101&fqt=1&beg=0&end=20500101&smplmt=460&lmt=1000000&_=1665632758059`
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	result, _ := io.ReadAll(res.Body)
	//resultString := string(result)
	reg := regexp.MustCompile(`^jQuery[\d_]+\(`)
	rst := reg.ReplaceAll(result, []byte{})
	rst = rst[:len(rst)-2]
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(rst, &resultMap)
	if err != nil {
		return "", "", ""
	}
	klines := resultMap["data"].(map[string]any)["klines"].([]any)
	split := strings.Split(klines[len(klines)-1].(string), ",")

	return split[2], split[9], split[8]
}
