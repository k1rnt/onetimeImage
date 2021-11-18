package config

import (
	"fmt"
	"os"

	"github.com/k1rnt/onetimeImage/domain/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewDB() *gorm.DB {
	connect := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), "tcp", os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))

	db, err := gorm.Open(os.Getenv("DBMS"), connect)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(model.Health{})

	return db
}
