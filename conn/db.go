package conn

import (
	"book-inventory/models"
	"log"
	"os"

	"github.com/lpernett/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal(errEnv.Error())
	}
	dsn := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("db not connected")
	}
	Migrate(db)
	return db
}

func Migrate(db *gorm.DB) {

	db.AutoMigrate(&models.Book{})

	data := models.Book{}
	if db.First(&data).Error == gorm.ErrRecordNotFound {
		seederBook(db)
	}

}

func seederBook(db *gorm.DB) {
	data := []models.Book{
		{
			Title:       "Perjalanan Mimpi",
			Author:      "Hasby Zain",
			Description: "Perjalanan di dalam mimpi",
			Stock:       10,
		},
		{
			Title:       "Manusia Serigala",
			Author:      "Sabiq Al",
			Description: "Cerita tentang manusia serigala",
			Stock:       2,
		},
	}

	for _, val := range data {
		db.Create(&val)
	}
}
