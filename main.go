package main

import (
	"bytes"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	_ "github.com/sijms/go-ora/v2"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	fmt.Println(`kiss zzy, put her legs on my shoulder, and lick her pretty foot, zzy, I want your body, your face`)
}

func startRead(ch chan string) {
	ch <- `big bang!`
}

func main_chan() {
	ch1 := make(chan string, 100)

	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- "第" + strconv.Itoa(i) + "个数"

		}
	}()

	go func() {
		for {
			received, ok := <-ch1
			fmt.Println("获取到了：", received, "ok:", ok)
			if strings.Contains(received, "78") {
				close(ch1)
			}
			if !ok {
				break
			}
		}
	}()

	time.Sleep(1 * time.Minute)
}

func testQuery(ctx *gin.Context) {
	conn, _ := sql.Open("oracle", "oracle://house_net:house_net@localhost:1521/orcl")
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			panic("关闭连接时出错:" + err.Error())
		}
	}(conn)
	rows, err := conn.Query(`select XH, JYZQC, FWZL, YWBJSJ 
									from trade_record where rownum <= 20
									`)
	if err != nil {
		println(err.Error())
	}
	var (
		row struct {
			XH     uint64
			JYZQC  string
			FWZL   string
			YWBJSJ string
		}
	)

	jsonResult := "["
	for rows.Next() {
		if err != rows.Scan(&row.XH, &row.JYZQC, &row.FWZL, &row.YWBJSJ) {
			fmt.Println("reading rows error occurred", err.Error())
			return
		}
		if marshal, err := json.Marshal(row); err == nil {
			s := string(marshal)
			jsonResult += s + ","
		} else {

		}
	}
	if jsonResult[len(jsonResult)-1:] == "," {
		jsonResult = jsonResult[:len(jsonResult)-1]
	}
	jsonResult += "]"
	ctx.String(http.StatusOK, jsonResult)
}

func name(ctx *gin.Context) {
	arr := map[string]any{
		"姓名": "兰州",
		"年龄": 888,
		"组织": nil,
	}
	ctx.JSON(http.StatusOK, arr)
}

func startListener(ctx *gin.Context) {
	port, b := ctx.GetQuery("p")
	if !b {
		port = strconv.Itoa(12345)
	}
	addr := net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 5000,
	}
	listener, err := net.ListenUDP("udp", &addr)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "socket created in fail.")
		return
	}
	ctx.Set("udp", listener)
	ctx.String(http.StatusOK, "socket has been created, port is %s", port)
}

type Animal int

const (
	Unknown Animal = iota
	Gopher
	Zebra
)

func (a *Animal) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	fmt.Println("it is ", a)
	fmt.Println("*it is ", *a)

	switch strings.ToLower(s) {
	default:
		*a = Unknown
	case "gopher":
		*a = Gopher
	case "zebra":
		*a = Zebra
	}

	return nil
}

type Person struct {
	Name string
	Age  uint
}

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
	var buf bytes.Buffer

	lista := []uint32{1, 2, 3, 4, 5}
	fmt.Println("lista is ", lista)
	fmt.Println("lista length is ", len(lista))
	fmt.Println("lista slice len is ", len(lista[:3]))
	fmt.Println("lista slice cap is ", cap(lista[:3]))

	listener, err := net.Listen("tcp", "0.0.0.0:4000")
	if err != nil {
		panic("监听创建失败,port:4000, caused by: " + err.Error())
	}

	fmt.Println("tcp listener created success.")
	defer listener.Close()
	for {
		client, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed.")
		}

		go func() {
			addr := client.RemoteAddr()
			fmt.Printf("accepted pack, network is: %s remote addr is: %s\n ------------------ stream below --------------------\n", addr.Network(), addr.String())
			//io.Copy(os.Stdout, client)
			io.Copy(&buf, client)
			io.Copy(os.Stdout, &buf)
		}()
	}
}

func main2() {
	err := DouBanBook()
	if err != nil {
		log.Fatalln(err.Error())
	}
	return

	//server := gin.Default()
	//server.GET("/name", name)
	//server.GET("/udp", startListener)
	//server.GET("/query", testQuery)
	//err := server.Run(":80")
	//if err != nil {
	//	return
	//}
}

var localAddr *string = flag.String("l", "localhost:9999", "local address")
var remoteAddr *string = flag.String("r", "localhost:80", "remote address")

func proxyConn(conn *net.TCPConn) {
	rAddr, err := net.ResolveTCPAddr("tcp", *remoteAddr)
	if err != nil {
		panic(err)
	}

	rConn, err := net.DialTCP("tcp", nil, rAddr)
	if err != nil {
		panic(err)
	}
	defer rConn.Close()

	buf := &bytes.Buffer{}
	for {
		data := make([]byte, 256)
		n, err := conn.Read(data)
		if err != nil {
			panic(err)
		}
		buf.Write(data[:n])
		if data[0] == 13 && data[1] == 10 {
			break
		}
	}

	if _, err := rConn.Write(buf.Bytes()); err != nil {
		panic(err)
	}
	log.Printf("sent:\n%v", hex.Dump(buf.Bytes()))

	data := make([]byte, 1024)
	n, err := rConn.Read(data)
	if err != nil {
		if err != io.EOF {
			panic(err)
		} else {
			log.Printf("received err: %v", err)
		}
	}
	log.Printf("received:\n%v", hex.Dump(data[:n]))
}

func handleConn(in <-chan *net.TCPConn, out chan<- *net.TCPConn) {
	for conn := range in {
		proxyConn(conn)
		out <- conn
	}
}

func closeConn(in <-chan *net.TCPConn) {
	for conn := range in {
		conn.Close()
	}
}

func main3() {

	flag.Parse()

	fmt.Printf("Listening: %v\nProxying: %v\n\n", *localAddr, *remoteAddr)

	addr, err := net.ResolveTCPAddr("tcp", *localAddr)
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}

	pending, complete := make(chan *net.TCPConn), make(chan *net.TCPConn)

	for i := 0; i < 5; i++ {
		go handleConn(pending, complete)
	}
	go closeConn(complete)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			panic(err)
		}
		pending <- conn
	}
}
