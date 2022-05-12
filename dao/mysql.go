package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:Lishoujia3301@tcp(120.27.194.228:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:root@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	//Migrate()
	return err
}

//func Migrate() {
//	DB.AutoMigrate(&model.User{})
//	DB.AutoMigrate(&model.Video{})
//	DB.AutoMigrate(&model.Comment{})
//	DB.AutoMigrate(&model.Like{})
//	DB.AutoMigrate(&model.Follow{})
//}
