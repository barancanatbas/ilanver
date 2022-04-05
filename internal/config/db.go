package config

import (
	"ilanver/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {

	// godotenv.Load(".env")

	// USER := os.Getenv("USER")
	// PASSWORD := os.Getenv("PASSWORD")
	// HOST := os.Getenv("HOST")
	// DBNAME := os.Getenv("DBNAME")

	dsn := "root:mysql123@tcp(127.0.0.1:3306)/dbilanver2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	DB = db

	Pool = newPool()
	//cleanupHook()
}

func Migrate() {
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&model.User{},
		&model.UserDetail{},
		&model.Adress{},
		&model.District{},
		&model.Province{},
		&model.Category{},
		&model.LostPassword{},
		&model.ProductDetail{},
		&model.ProductState{},
		&model.Product{},
		&model.Promo{},
	)
}
