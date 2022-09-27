package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/sijms/go-ora/v2"
)

func testQuery() {
	conn, err := sql.Open("oracle", "oracle://house_net:house_net@localhost:1521/orcl")
	if err != nil {
		fmt.Println("数据连接异常:", err.Error())
		return
	} else {
		err := conn.Ping()
		if err != nil {
			fmt.Println("数据连接异常:", err.Error())
			return
		} else {
			fmt.Println("数据连接成功:")
		}
	}
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			panic("关闭连接时出错:" + err.Error())
		}
	}(conn)
	rows, err := conn.Query(`select XH, JYZQC, FWZL, YWBJSJ from trade_record where rownum <= 20 `)
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
		err := rows.Scan(&row.XH, &row.JYZQC, &row.FWZL, &row.YWBJSJ)
		if err != nil {
			fmt.Println("reading rows error occurred", err.Error())
			return
		} else {

			if marshal, err := json.Marshal(row); err == nil {
				s := string(marshal)
				jsonResult += s + ","
			} else {

			}
		}
	}
	if jsonResult[len(jsonResult)-1:] == "," {
		jsonResult = jsonResult[:len(jsonResult)-1]
	}
	jsonResult += "]"
	fmt.Println(jsonResult)
}

func main() {
	testQuery()
}
