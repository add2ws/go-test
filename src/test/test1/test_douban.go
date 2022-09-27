package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strings"
)

// 豆瓣书榜单
func DouBanBook() error {
	// 创建 Collector 对象
	collector := colly.NewCollector()
	// 在请求之前调用
	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("回调函数OnRequest: 在请求之前调用")
	})
	// 请求期间发生错误,则调用
	collector.OnError(func(response *colly.Response, err error) {
		fmt.Println("回调函数OnError: 请求错误", err)
	})
	// 收到响应后调用
	collector.OnResponse(func(response *colly.Response) {
		fmt.Println("回调函数OnResponse: 收到响应后调用")
	})
	//OnResponse如果收到的内容是HTML ,则在之后调用
	collector.OnHTML("ul[class='subject-list']", func(element *colly.HTMLElement) {
		// 遍历li
		element.ForEach("li", func(i int, el *colly.HTMLElement) {
			// 获取封面图片
			coverImg := el.ChildAttr("div[class='pic'] > a[class='nbg'] > img", "src")
			// 获取书名
			bookName := el.ChildText("div[class='info'] > h2")
			// 获取发版信息，并从中解析出作者名称
			authorInfo := el.ChildText("div[class='info'] > div[class='pub']")
			split := strings.Split(authorInfo, "/")
			author := split[0]
			fmt.Printf("封面: %v 书名:%v 作者:%v\n", coverImg, trimSpace(bookName), author)
		})
	})
	// 发起请求
	return collector.Visit("https://book.douban.com/tag/小说")
}

// 删除字符串中的空格信息
func trimSpace(str string) string {
	// 替换所有的空格
	str = strings.ReplaceAll(str, " ", "")
	// 替换所有的换行
	return strings.ReplaceAll(str, "\n", "")
}

func main() {
	err := DouBanBook()
	if err != nil {
		log.Fatalln(err.Error())
	}

}
