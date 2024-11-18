package dbconnection

import (
	"fmt"

	"github.com/handarudwiki/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabaseConnection(config *config.Config) (*gorm.DB, error) {
	fmt.Println(config)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Database.Host, config.Database.User, config.Database.Password, config.Database.Name, config.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
