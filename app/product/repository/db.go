package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/FelizYK/gomall/app/product/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	// init MySQL
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt: true,
		},
	)
	if err != nil {
		panic(err)
	}

	// demo
	if os.Getenv("GO_ENV") != "online" && !DB.Migrator().HasTable(&Product{}) {
		log.Printf("Insert demo product data ...")
		DB.AutoMigrate(&Product{}, &Category{})
		DB.Exec("INSERT INTO `gomall`.`category` VALUES (1,'2023-12-06 15:05:06','2023-12-06 15:05:06',NULL,'T-Shirt','T-Shirt'), (2,'2023-12-06 15:05:06','2023-12-06 15:05:06',NULL,'Sticker','Sticker')")
		DB.Exec("INSERT INTO `gomall`.`product` VALUES ( 1, '2023-12-06 15:26:19', '2023-12-09 22:29:10', NULL, 'Notebook', 'The cloudwego notebook is a highly efficient and feature-rich notebook designed to meet all your note-taking needs. ', '/assets/notebook.jpeg', 9.90 ), ( 2, '2023-12-06 15:26:19', '2023-12-09 22:29:10', NULL, 'Mouse-Pad', 'The cloudwego mouse pad is a premium-grade accessory designed to enhance your computer usage experience. ', '/assets/mouse-pad.jpeg', 8.80 ), ( 3, '2023-12-06 15:26:19', '2023-12-09 22:31:20', NULL, 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/assets/t-shirt-1.jpeg', 6.60 ), ( 4, '2023-12-06 15:26:19', '2023-12-09 22:31:20', NULL, 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/assets/t-shirt-2.jpeg', 2.20 ), ( 5, '2023-12-06 15:26:19', '2023-12-09 22:32:35', NULL, 'Sweatshirt', 'The cloudwego Sweatshirt is a cozy and fashionable garment that provides warmth and style during colder weather.', '/assets/sweatshirt.jpeg', 1.10 ), ( 6, '2023-12-06 15:26:19', '2023-12-09 22:31:20', NULL, 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/assets/t-shirt-3.jpeg', 1.80 ), ( 7, '2023-12-06 15:26:19', '2023-12-09 22:31:20', NULL, 'mascot', 'The cloudwego mascot is a charming and captivating representation of the brand, designed to bring joy and a playful spirit to any environment.', '/assets/logo.jpg', 4.80 )")
		DB.Exec("INSERT INTO `gomall`.`product_category` (product_id,category_id) VALUES ( 1, 2 ), ( 2, 2 ), ( 3, 1 ), ( 4, 1 ), ( 5, 1 ), ( 6, 1 ),( 7, 2 )")
	}
}
