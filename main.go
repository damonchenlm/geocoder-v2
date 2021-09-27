package main

import (
	"geocoder-v2/global"
	"geocoder-v2/model"
	"geocoder-v2/router"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//数据库相关， 暂时写在 main 内
	var err error
	global.DB, err = gorm.Open("mysql", "root:Cyl851106@(127.0.0.1:3306)/geocoder?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer global.DB.Close()

	// 自动迁移
	global.DB.AutoMigrate(&model.Location{})

	engine := router.SetupRouter()
	_ = engine.Run()
}
