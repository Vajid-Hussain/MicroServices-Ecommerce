package db_product_svc

import (
	"database/sql"
	"fmt"

	config_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config *config_product_svc.Config) (*gorm.DB, error) {
	connectionString := "user=postgres password=123 host=localhost"
	sql, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	rows, err := sql.Query("SELECT 1 FROM pg_database WHERE datname = "+ config.DBName)
	if err != nil {
		fmt.Println("Error checking database existence:", err)
	}
	defer rows.Close()

	// If the database exists, do nothing
	if rows.Next() {
		fmt.Println("Database"+config.DBName+" already exists.")
	} else {
		// If the database does not exist, create it
		_, err = sql.Exec("CREATE DATABASE " + config.DBName)
		if err != nil {
			fmt.Println("Error creating database:", err)
		}
	}

	DB, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// err = DB.AutoMigrate(domain_auth_svc.Users{},
	// 	domain_auth_svc.Admin{})
	// if err != nil {
	// 	return nil, err
	// }

	return DB, nil
}
