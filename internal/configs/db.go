package config

import (
	"ilanver/internal/models"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

// bu nesneyi model kısmında çağırırım
var Database *gorm.DB

// vt ayarlarını çalıştırma methodu
func Init() {
	Database = Connect()
	AutoMigrate()
}

// bağlantı açar
func Connect() *gorm.DB {
	godotenv.Load(".env")

	USER := os.Getenv("USER")
	PASSWORD := os.Getenv("PASSWORD")
	HOST := os.Getenv("HOST")
	DBNAME := os.Getenv("DBNAME")

	db, err := gorm.Open("mysql", USER+":"+PASSWORD+HOST+DBNAME+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err.Error())
	}
	return db
}

// otomatik migrate yapar
func AutoMigrate() *gorm.DB {
	migrate := Database.AutoMigrate(
		&models.Category{},
		&models.Photo{},
		&models.Promo{},
		&models.PromoRequest{},
		&models.Province{},
		&models.District{},
		&models.Adress{},
		&models.UserDetail{},
		&models.User{},
		&models.LostPassword{},
		&models.Product{},
		&models.ProductState{},
		&models.ProductDetail{},
	)

	// tbl photos = resimlerin tutulduğu tablo
	Database.Model(&models.Photo{}).AddForeignKey("productfk", "products(id)", "cascade", "cascade")

	// tbl promo = reklamların tablosu
	Database.Model(&models.Promo{}).AddForeignKey("categoryfk", "categories(id)", "cascade", "cascade")
	Database.Model(&models.Promo{}).AddForeignKey("photofk", "photos(id)", "cascade", "cascade")

	// tbl promo request = gelen reklam istekleri
	Database.Model(&models.PromoRequest{}).AddForeignKey("categoryfk", "categories(id)", "cascade", "cascade")
	Database.Model(&models.PromoRequest{}).AddForeignKey("photofk", "photos(id)", "cascade", "cascade")

	// tbl district = ilçelerin tabloları
	Database.Model(&models.District{}).AddForeignKey("provincefk", "provinces(id)", "cascade", "cascade")

	// tbl adress = adres ile ilgili bilgileri tutar
	Database.Model(&models.Adress{}).AddForeignKey("districtfk", "districts(id)", "cascade", "cascade")

	// tbl user detail = user bilgilerinin detaylarını tutar
	Database.Model(&models.UserDetail{}).AddForeignKey("adressfk", "adresses(id)", "cascade", "cascade")

	// tbl user = user bilgilerini tutar
	Database.Model(&models.User{}).AddForeignKey("userdetailfk", "user_details(id)", "cascade", "cascade")

	// tbl lostpassword = şifremi unuttum alanı ile ilgili
	Database.Model(&models.LostPassword{}).AddForeignKey("userfk", "users(id)", "cascade", "cascade")

	// tbl product
	Database.Model(&models.Product{}).AddForeignKey("userfk", "users(id)", "cascade", "cascade")

	// product detail
	Database.Model(&models.ProductDetail{}).AddForeignKey("adressfk", "adresses(id)", "cascade", "cascade")
	Database.Model(&models.ProductDetail{}).AddForeignKey("categoryfk", "categories(id)", "cascade", "cascade")

	return migrate
}
