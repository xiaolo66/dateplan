package model

import (
	"dateplan/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

var (
	DB *gorm.DB
	Once sync.Once
	err  error
)

func InitDB(){
	Once.Do(func() {
		DB,err=gorm.Open("mysql","root:Monika8899174!@tcp(127.0.0.1:3306)/Bubble?charset=utf8mb4&parseTime=True&loc=Local")
		if err!=nil{
			utils.Log.Errorf("open DB failed,err:",err)
			return
		}
		DB.AutoMigrate(&Bubble{})
		utils.Log.Info("Init Db success")
	})
}
