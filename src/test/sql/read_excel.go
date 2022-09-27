package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

var allTables [][]string
var allNeedTables []string

func GetByTable(tbName string) []string {
	if allTables == nil {
		allTables, allNeedTables, _ = ReadExcel()
	}

	var results []string
	tbNameFull := "U_SJZX_ODSK." + tbName
	//fmt.Println("开始根据表名过滤...", tbNameFull)
	for _, tbAndCol := range allTables {
		if tbAndCol[0] == tbNameFull {
			results = append(results, tbAndCol[1])
		}
	}
	return results
}

func GetAllNeedTables() []string {
	if allTables == nil {
		allTables, allNeedTables, _ = ReadExcel()
	}
	return allNeedTables
}

func ReadExcel() ([][]string, []string, error) {
	excel, err := excelize.OpenFile("D:/白皮书数据最新（标注版）.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	defer func() {
		if err := excel.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := excel.GetRows("SQL Results")
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	var list [][]string
	for _, row := range rows {
		list = append(list, []string{row[3], row[4]})
	}

	rows2, err := excel.GetRows("SQL Statement")
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	var list2 []string
	for _, row := range rows2 {
		list2 = append(list2, row[3])
	}
	return list, list2, nil
}
