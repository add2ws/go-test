package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/sijms/go-ora/v2"
	"io"
	"net"
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

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:4000")
	if err != nil {
		panic("监听创建失败,port:4000, caused by: " + err.Error())
	}

	fmt.Println("tcp listener created success.")
	defer listener.Close()
	i := 0
	for {
		client, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed.")
		} else {
			i++
			fmt.Println("accept client...", i)
		}

		go func() {
			addr := client.RemoteAddr()
			fmt.Printf("accepted pack, network is: %s remote addr is: %s\n ------------------ stream below --------------------\n", addr.Network(), addr.String())
			//io.Copy(os.Stdout, client)
			io.Copy(os.Stdout, client)
		}()
	}
}
