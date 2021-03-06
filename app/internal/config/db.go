package config

import (
	"ilanver/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB, DBTest *gorm.DB
var ElasticDB *ElasticSearch

func Init() {

	dsn := "root:mysql123@tcp(127.0.0.1:3306)/dbilanver2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	DB = db

	// redis config
	Pool = NewPool()

	// elastic config
	ElasticDB, err = NewElastic([]string{"http://localhost:9200"})

}

func InitTest() {
	dsn := "root:mysql123@tcp(127.0.0.1:3306)/ilanverdb-test?charset=utf8mb4&parseTime=True&loc=Local"
	dbTest, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	DBTest = dbTest

	// redis config
	Pool = NewPool()

	// elastic config
	ElasticDB, err = NewElastic([]string{"http://localhost:9200"})
	//Migrate(DBTest)
}

func Migrate(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
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
