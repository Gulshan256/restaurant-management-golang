package initializers

import (
	// "gorm.io/driver/postgres"
	"github.com/devGulshan/restaurant-management/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {

	var err error
	// DB, err = gorm.Open(postgres.Open("postgres://postgres:123@localhost:5432/go_tut"), &gorm.Config{})
	DB, err = gorm.Open(sqlite.Open("foodCollection.db"), &gorm.Config{})
	if err != nil {
		panic(" to connect database failed")
	}
	DB.AutoMigrate(&models.Food{}, &models.Invoice{}, &models.Menu{}, &models.Order{}, &models.OrderItem{}, &models.Table{}, &models.User{}, &models.Note{})

}
