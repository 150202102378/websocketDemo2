package dao

import (
	"log"
	"os"
	"websocketDemo2/conf"

	"github.com/jinzhu/gorm"
)

var (
	psqlDB *gorm.DB
)

func init() {
	initTable(Record{})
}

//创建table
func initTable(structData interface{}) {
	if psqlDB == nil {
		psqlDB = conf.GetPsqlDB()
	}
	if !psqlDB.HasTable(structData) {
		if err := psqlDB.CreateTable(structData).Error; err != nil {
			log.Fatal("init table error")
			os.Exit(3)
		}
	}
}
