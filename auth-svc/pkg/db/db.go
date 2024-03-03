package db_auth_svc

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	config_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/config"
	domain_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/domain"
	helper_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config *config_auth_svc.General) (*gorm.DB, error) {
	connectionString := "user=postgres password=123 host=localhost"
	sql, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// _, err = sql.Exec("CREATE DATABASE " + config.DBName)
	// if err != nil {
	// 	fmt.Println("err ", err)
	// 	return nil, err
	// }

	rows, err := sql.Query("SELECT 1 FROM pg_database WHERE datname = 'auth_svc_ecommerce'")
	if err != nil {
		fmt.Println("Error checking database existence:", err)
	}
	defer rows.Close()

	// If the database exists, do nothing
	if rows.Next() {
		fmt.Println("Database 'auth_svc_ecommerce' already exists.")
	} else {
		// If the database does not exist, create it
		_, err = sql.Exec("CREATE DATABASE auth_svc_ecommerce")
		if err != nil {
			fmt.Println("Error creating database:", err)
		}
	}

	DB, err := gorm.Open(postgres.Open(config.DbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = DB.AutoMigrate(domain_auth_svc.Users{},
		domain_auth_svc.Admin{})
	if err != nil {
		return nil, err
	}

	CheckAndCreateAdmin(DB)

	return DB, nil
}

func CheckAndCreateAdmin(DB *gorm.DB) {
	var count int
	var (
		Name     = "mobileMart"
		Email    = "mobilemart@gmail.com"
		Password = "buyMobiles"
	)
	HashedPassword := helper_auth_svc.HashPassword(Password)

	query := "SELECT COUNT(*) FROM admins"
	DB.Raw(query).Row().Scan(&count)
	if count <= 0 {
		query = "INSERT INTO admins(name, email, password) VALUES(?, ?, ?)"
		DB.Exec(query, Name, Email, HashedPassword).Row().Err()
	}
}
