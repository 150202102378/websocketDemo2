package dao

import (
	"fmt"
)

//记录数据，获取数据成功记录为1，失败为0
func RecordMessage(record *Record) bool {
	tx := psqlDB.Begin()
	err := psqlDB.Create(record).Error
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}
