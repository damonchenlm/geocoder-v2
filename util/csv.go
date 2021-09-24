package util

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCsv(filepath string) [][]string {
	//打开文件(只读模式)，创建io.read接口实例
	opencast, err := os.Open(filepath)
	if err != nil {
		log.Println("csv文件打开失败！")
	}
	defer opencast.Close()

	//创建csv读取接口实例
	readCsv := csv.NewReader(opencast)

	//读取所有内容
	readAll, err := readCsv.ReadAll() //返回切片类型：[[s s ds] [a a a]]
	return readAll

}
