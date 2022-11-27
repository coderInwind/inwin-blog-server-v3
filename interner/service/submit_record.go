package service

import (
	"fmt"
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/interner/model"
	"time"
)

type SubmitRecordService struct {
}

func (*SubmitRecordService) CreateSubmitRecord(count int) {
	if count == 0 {
		return
	}
	// 查询是否已有该行
	record := model.SubmitRecord{}
	newDate := time.Now().Format("2006-01-02")
	err := global.DB.First(&record, "date = ?", newDate).Error
	if err != nil {
		//没有就创建
		newRecord := model.SubmitRecord{
			Date:  time.Now(),
			Count: count, //获取utf-8长度
		}
		err := global.DB.Create(&newRecord).Error
		fmt.Println(err)
	} else {
		oldCount := record.Count
		//有则将当前的字数加上已经有的字数
		global.DB.Model(&model.SubmitRecord{}).Where("date = ?", newDate).Update("count", oldCount+count)
	}
}

func (*SubmitRecordService) GetSubmitRecord(date []string) (records []model.SubmitRecord, err error) {

	//global.DB
	fmt.Println(date[0] < date[1])
	fmt.Println(date[0], date[1])
	err = global.DB.Where("date BETWEEN ? AND ?", date[0], date[1]).Find(&records).Error
	return records, err
}
