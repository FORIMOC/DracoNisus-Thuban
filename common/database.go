package common

import (
	"fmt"

	"forimoc.DracoNisus-Thuban/model/ModelAdmin"
	"forimoc.DracoNisus-Thuban/model/ModelGroup"
	"forimoc.DracoNisus-Thuban/model/ModelKeyUser"
	"forimoc.DracoNisus-Thuban/model/ModelKeyword"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 保存数据库实例的全局变量
var DB *gorm.DB

// InitDB 数据库配置初始化
// =>
func InitDB() {
	// 从配置文件 application.yml 获得配置信息
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	// loc=Local 设置本地时区
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		LoggerTer.Fatal("connect database failed: " + err.Error())
	}
	// 自动将模型迁移至数据库表
	err = db.AutoMigrate(
		&ModelAdmin.Admin{},
		&ModelGroup.Group{},
		&ModelGroup.GroupCacheMessage{},
		&ModelGroup.GroupGraph{},
		&ModelGroup.GroupKeyword{},
		&ModelGroup.GroupMessage{},
		&ModelGroup.GroupTime{},
		&ModelGroup.GroupUserTime{},
		&ModelGroup.GroupUser{},
		&ModelKeyword.KeywordList{},
		&ModelKeyUser.KeyUser{},
		&ModelKeyUser.KeyUserTime{},
		&ModelKeyUser.KeyUserCacheMessage{},
		&ModelKeyUser.KeyUserKeyword{},
		&ModelKeyUser.KeyUserGroup{},
	)
	if err != nil {
		LoggerTer.Fatal("automigrate model failed: " + err.Error())
	}
	DB = db
}

// GetDB 返回数据库实例
// => 数据库实例
func GetDB() *gorm.DB {
	return DB
}
