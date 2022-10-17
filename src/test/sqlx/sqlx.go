package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/sijms/go-ora/v2"
)

func testSqlx() error {
	db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	} else {
		fmt.Println("数据库连接成功！")
	}

	db.Query("select * from trade_record")

	return nil
}

func main() {
	testSqlx()
}
