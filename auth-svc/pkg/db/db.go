package db_auth_svc

import (
	"database/sql"

	config_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/config"
	domain_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config *config_auth_svc.General) (*gorm.DB, error) {
	connectionString := "user= vajid password=5432 host= localhost"
	sql, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	_, err = sql.Exec("create database " + config.DBName)
	if err != nil {
		return nil, err
	}

	DB, err := gorm.Open(postgres.Open(config.DbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = DB.AutoMigrate(domain_auth_svc.Users{})
	if err != nil {
		return nil, err
	}

	return DB, nil
}
