package model

import (
	"dateplan/config"
	"dateplan/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

var (
	DB *gorm.DB
	Once sync.Once
	err  error
)

func InitDB(cfg config.Config){
	Once.Do(func() {
		s:=fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",cfg.Mysql.Username,cfg.Mysql.Password,cfg.Mysql.Addr,cfg.Mysql.DbName)
		fmt.Println(s)
		DB,err=gorm.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",cfg.Mysql.Username,cfg.Mysql.Password,cfg.Mysql.Addr,cfg.Mysql.DbName))
		if err!=nil{
			utils.Log.Errorf("open DB failed,err:",err)
			return
		}
		DB.AutoMigrate(&Bubble{})
		utils.Log.Info("Init Db success")
	})
}
