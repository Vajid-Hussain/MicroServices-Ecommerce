package db_order_svc

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	config_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/config"
	domain_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config *config_order_svc.Config) (*gorm.DB, error) {
	connectionString := "user=postgres password=123 host=localhost"
	sql, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	defer sql.Close()

	// Check if the database exists
	rows, err := sql.Query("SELECT 1 FROM pg_database WHERE datname = $1", config.DBName)
	if err != nil {
		fmt.Println("Error checking database existence:", err)
		return nil, err
	}
	defer rows.Close()

	// If the database exists, do nothing
	if rows.Next() {
		fmt.Println("Database", config.DBName, "already exists.")
	} else {
		// If the database does not exist, create it
		_, err = sql.Exec("CREATE DATABASE " + config.DBName)
		if err != nil {
			fmt.Println("Error creating database:", err)
			return nil, err
		}
		fmt.Println("Database", config.DBName, "created.")
	}

	// Connect to the specified database
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.DBUrl,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&domain_order_svc.Order{}, &domain_order_svc.OrderProducts{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// func InitDB(config *config_order_svc.Config) (*gorm.DB, error) {
// 	fmt.Println("@@@", config)
// 	connectionString := "user=postgres password=123 host=localhost"
// 	sql, err := sql.Open("postgres", connectionString)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rows, err := sql.Query("SELECT 1 FROM pg_database WHERE datname = '" + config.DBName +"'")
// 	if err != nil {
// 		fmt.Println("Error checking database existence:", err)
// 	}
// 	defer rows.Close()

// 	// If the database exists, do nothing
// 	if rows.Next() {
// 		fmt.Println("Database" + config.DBName + " already exists.")
// 	} else {
// 		// If the database does not exist, create it
// 		_, err = sql.Exec("CREATE DATABASE '" + config.DBName +"'")
// 		if err != nil {
// 			fmt.Println("Error creating database:", err)
// 		}
// 	}

// 	DB, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = DB.AutoMigrate(domain_order_svc.Order{},
// 		domain_order_svc.OrderProducts{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return DB, nil
// }
