package data

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strings"
)

// PasswordItem 读取原数据的每一次条目
type PasswordItem struct {
	Category string
	Account  string
	Password string
	Site     string
}

func GetData(filepath string) ([]PasswordItem, error) {
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		//log.Println(err)
		return nil, err
	}

	res := make([]PasswordItem, 0)
	//
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		//log.Println(err)
		return nil, err
	}

	//log.Printf("rows.len is %v", len(rows))
	//  计算索引

	cellList := []string{"A", "B", "C", "D"}
	for index := 2; index <= len(rows); index++ {
		passwordItem := make([]string, len(cellList))

		for cellIndex, cell := range cellList {
			value, err := f.GetCellValue("Sheet1", fmt.Sprintf("%s%d", cell, index))
			if err != nil {
				//log.Println("获取值发生错误,", err)
				value = ""
			}
			//fmt.Println("value is ", value)
			passwordItem[cellIndex] = strings.TrimSpace(value)
		}
		res = append(res, PasswordItem{
			Category: passwordItem[0],
			Account:  passwordItem[1],
			Password: passwordItem[2],
			Site:     passwordItem[3],
		})
	}

	//for i := 0; i < len(res); i++ {
	//	log.Printf("%+v\n", res[i])
	//}

	return res, nil
}
