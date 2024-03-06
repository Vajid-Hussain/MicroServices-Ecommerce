package db_order_svc

import (
	config_order_svc "command-line-arguments/home/vajid/Brocamp/Mobile-Mart-Microservice/order-svc/pkg/config/config.go"
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config *config_order_svc.Config) (*gorm.DB, error) {
	connectionString := "user=postgres password=123 host=localhost"
	sql, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	rows, err := sql.Query("SELECT 1 FROM pg_database WHERE datname = '" + config.DBName + "'")
	if err != nil {
		fmt.Println("Error checking database existence:", err)
	}
	defer rows.Close()

	// If the database exists, do nothing
	if rows.Next() {
		fmt.Println("Database" + config.DBName + " already exists.")
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

	// err = DB.AutoMigrate(domain_products_svc.Brand{},
	// 	domain_products_svc.Category{},
	// 	domain_products_svc.Inventories{},
	// 	domain_products_svc.Cart{})
	// if err != nil {
	// 	return nil, err
	// }

	return DB, nil
}
