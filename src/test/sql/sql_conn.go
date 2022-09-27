package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/sijms/go-ora/v2"
	"sort"
	"strings"
)

var conn *sql.DB = nil

func main() {
	tables := GetAllNeedTables()

	for i, table := range tables {
		tb := strings.ReplaceAll(table, "U_SJZX_ODSK.", "")
		str := compareTable(tb, false)
		println(i+1, tb, str)
	}
	conn.Close()
}

func main2() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		panic("没有传入表名")
		return
	}
	tableName := args[0]
	fmt.Println("表名为：", tableName)

	compareTable(tableName, true)
	conn.Close()
}

func compareTable(tableName string, print bool) string {
	colFromExcel := GetByTable(tableName)
	sort.Strings(colFromExcel)

	colFromDB := readDB(tableName)
	sort.Strings(colFromDB)

	minLen := len(colFromDB)
	if len(colFromDB) > len(colFromExcel) {
		minLen = len(colFromExcel)
	}

	for i := 0; i < minLen; i++ {
		if print {
			fmt.Println(i+1, colFromExcel[i], colFromDB[i])
		}
	}

	if len(colFromDB) == 0 {
		if print {
			fmt.Println("表在数据库里不存在", tableName)
		}
		return "表在数据库里不存在"
	}

	s1 := notIn(colFromExcel, colFromDB)
	s2 := notIn(colFromDB, colFromExcel)

	if len(s1) == 0 && len(s2) == 0 {
		if print {
			fmt.Println("一致")
		}
		return "一致"
	}

	var results string
	results += "excel里多的=======>"
	if print {
		fmt.Println("excel里多的=======>")
	}
	for _, s := range s1 {
		results += s + " "
		if print {
			fmt.Println(s)
		}
	}
	results += "数据库里多的=======>"
	if print {
		fmt.Println("数据库里多的=======>")
	}
	for _, s := range s2 {
		results += s
		if print {
			fmt.Println(s)
		}
	}

	return results
}

func notIn(s1 []string, s2 []string) []string {
	var results []string
	for _, s11 := range s1 {
		exist := false
		for _, s22 := range s2 {
			if s11 == s22 {
				exist = true
				continue
			}
		}
		if !exist {
			results = append(results, s11)
		}
	}
	return results
}

func readDB(tableName string) []string {
	//conn, err := sqlStr.Open("oracle", "oracle://u_jw_sjdj:jw123456@tldata:1522/orcl")
	if conn == nil {
		var err error
		conn, err = sql.Open("oracle", "oracle://u_jw_sjdj:jw123456@172.16.1.220:1522/orcl")
		if err != nil {
			fmt.Println("数据库连接异常！", err.Error())
			return nil
		} else {
			fmt.Println("开始连接数据库......")
			err := conn.Ping()
			if err != nil {
				fmt.Println("数据库连接异常！", err.Error())
			} else {
				fmt.Println("数据库连接成功！")
			}
		}
	}

	sqlStr := "SELECT TABLE_NAME, COLUMN_NAME FROM ALL_TAB_COLUMNS  WHERE TABLE_NAME = '" + tableName + "'"
	rows, err := conn.Query(sqlStr)
	if err != nil {
		println("查询出错：" + err.Error())
	}
	defer rows.Close()

	row := struct {
		TABLE_NAME  string
		COLUMN_NAME string
	}{}

	var results []string
	for rows.Next() {
		err := rows.Scan(&row.TABLE_NAME, &row.COLUMN_NAME)
		if err != nil {
			fmt.Println("scan error:", err.Error())
		} else {
			if row.COLUMN_NAME != "" {
				results = append(results, row.COLUMN_NAME)
			} else {
				fmt.Println("检测到空字符串列")
			}
		}
	}
	return results
}
