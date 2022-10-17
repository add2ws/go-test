package main

import (
	"go-test/src/service"
	crawler2 "go-test/src/service/crawler"
	"io"
	"net/http"
	"time"
)

func getByStock(stockCode string) {
	url := `https://push2.eastmoney.com/api/qt/stock/trends2/get?secid=1.` + stockCode + `&fields1=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f11,f12,f13&fields2=f51,f52,f53,f54,f55,f56,f57,f58&ut=fa5fd1943c7b386f172d6893dbfba10b&iscr=0&cb=cb_1665645217100_48029913&isqhquote=&cb_1665645217100_48029913=cb_1665645217100_48029913`
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	result, _ := io.ReadAll(resp.Body)
	println(string(result))
}

func infoWithPrint(name string, fn func() (string, string, string)) {
	price, change, percent := fn()
	println(name+":", "price", price, "change=", change, "percent=", percent)
}

func infoSave() {
	redisCache := service.RedisCache

	price, change, percent := crawler2.BitcoinToCNH()
	redisCache.Set("bitcoin", price+";"+change+";"+percent)

	price, change, percent = crawler2.Shangzheng()
	redisCache.Set("shangzheng", price+";"+change+";"+percent)
	price, change, percent = crawler2.USDToCNH()
	redisCache.Set("usdtocnh", price+";"+change+";"+percent)

}

func showPrice() {
	for i := 0; i < 100; i++ {
		infoWithPrint("BTC/USD", crawler2.BitcoinToCNH)
		infoWithPrint("上证指数", crawler2.Shangzheng)
		infoWithPrint("USD/CNH", crawler2.USDToCNH)
		println("---------------------------------------------------------------")
		//go func() {
		//}()
		time.Sleep(time.Millisecond * 1000)
	}
}

func testCache() {
	redisCache := service.RedisCache

	println("setting...")
	go redisCache.Set("name", "王丽丽")

	time.Sleep(time.Second * 2)
	val := redisCache.GetString("nppame")
	println("val=>", val)
}

func testRedis() {
	for i := 0; i < 100; i++ {
		val := service.RedisCache.Incr("age")
		println("val=>", val)
	}
}

func getCurrentTime() string {
	now := time.Now()
	return now.Format("")
}

func main() {
	a, b, c := crawler2.USDToCNH()
	println(a, b, c)
}
