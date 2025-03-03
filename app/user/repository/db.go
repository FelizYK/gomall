package repository

import (
	"fmt"
	"os"

	"github.com/FelizYK/gomall/app/user/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	// init MySQL
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	db, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt: true,
		},
	)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
}
