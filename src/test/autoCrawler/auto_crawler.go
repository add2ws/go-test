package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"go-test/src/entity"
	"go-test/src/service"
	"go-test/src/service/crawler"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var usdList []entity.USDToCNH

func fetchUSDToCNH() {
	price, change, percent := crawler.USDToCNH()
	price = strings.ReplaceAll(price, ",", "")
	priceFloat, err := strconv.ParseFloat(price, 10)
	if err != nil {
		println(price, err.Error())
		return
	}
	dataRow := entity.NewUSDToCNH(float32(priceFloat), change, percent)
	if len(usdList) >= 100 {
		usdList = usdList[1 : len(usdList)-1]
	}
	usdList = append(usdList, dataRow)

	var lastItem entity.USDToCNH
	service.DBConn.Last(&lastItem)
	if lastItem.Price != float32(priceFloat) {
		service.DBConn.Create(&dataRow)
	}
}

var bitcoinList []entity.BitcoinToUSD

func fetchBitcoin() {
	price, change, percent := crawler.BitcoinToCNH()
	price = strings.ReplaceAll(price, ",", "")
	priceFloat, err := strconv.ParseFloat(price, 10)
	if err != nil {
		println(price, err.Error())
		return
	}
	dataRow := entity.NewBitcoinToUSD(float32(priceFloat), change, percent)
	if len(bitcoinList) >= 100 {
		bitcoinList = bitcoinList[1 : len(bitcoinList)-1]
	}
	bitcoinList = append(bitcoinList, dataRow)

	var lastItem entity.BitcoinToUSD
	service.DBConn.Last(&lastItem)
	if lastItem.Price != float32(priceFloat) {
		service.DBConn.Create(&dataRow)
	}
}

var shangzhengList []entity.Shangzheng

func fetchShangzheng() {
	price, change, percent := crawler.Shangzheng()
	price = strings.ReplaceAll(price, ",", "")
	priceFloat, err := strconv.ParseFloat(price, 10)
	if err != nil {
		println(price, err.Error())
		return
	}
	dataRow := entity.NewShangzheng(float32(priceFloat), change, percent)
	if len(shangzhengList) >= 100 {
		shangzhengList = shangzhengList[1 : len(shangzhengList)-1]
	}
	shangzhengList = append(shangzhengList, dataRow)

	var lastItem entity.Shangzheng
	service.DBConn.Last(&lastItem)
	if lastItem.Price != float32(priceFloat) {
		service.DBConn.Create(&dataRow)
	}
}

type DataItem struct {
	Datetime string
	Price    float32
	Change   string
	Percent  string
}

func getUSDToCNH(c *gin.Context) {
	var result []DataItem
	for _, item := range usdList {
		result = append(result, DataItem{
			Datetime: time.UnixMilli(item.ID).Format("2006/01/02 15:04:05"),
			Price:    item.Price,
			Change:   item.Change,
			Percent:  item.Percent,
		})
	}
	c.JSON(http.StatusOK, result)
}

func getBitcoinUSD(c *gin.Context) {
	var result []DataItem
	for _, item := range bitcoinList {
		result = append(result, DataItem{
			Datetime: time.UnixMilli(item.ID).Format("2006/01/02 15:04:05"),
			Price:    item.Price,
			Change:   item.Change,
			Percent:  item.Percent,
		})
	}
	c.JSON(http.StatusOK, result)
}

func getShangzheng(c *gin.Context) {
	var result []DataItem
	for _, item := range shangzhengList {
		result = append(result, DataItem{
			Datetime: time.UnixMilli(item.ID).Format("2006/01/02 15:04:05"),
			Price:    item.Price,
			Change:   item.Change,
			Percent:  item.Percent,
		})
	}
	c.JSON(http.StatusOK, result)
}

func main() {
	scheduler := gocron.NewScheduler(time.UTC)

	scheduler.Every(1).Second().Do(func() {
		fetchUSDToCNH()
	})
	scheduler.Every(1).Second().Do(fetchBitcoin)
	scheduler.Every(1).Second().Do(fetchShangzheng)

	scheduler.StartAsync()

	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.GET("/usdtocnh", getUSDToCNH)
	r.GET("/bitcoin", getBitcoinUSD)
	r.GET("/shangzheng", getShangzheng)
	r.Run(":7788")

}
