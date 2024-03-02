package repository_auth_svc

import (
	"fmt"

	requestmodel_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/model/requestmodel"
	responsemodel_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/model/responsemodel"
	interfaceRepository_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/repository/interface"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaceRepository_auth_svc.IuserRepository {
	return &userRepository{DB: db}
}

func (d *userRepository) IsUserExist(phone string) int {
	var userCount int

	query := "SELECT COUNT(*) FROM users WHERE phone=$1 AND status!=$2"
	err := d.DB.Raw(query, phone, "delete").Row().Scan(&userCount)
	if err != nil {
		fmt.Println("Error for user exist, using same phone in signup")
	}
	return userCount
}

func (d *userRepository) CreateUser(userDetails *requestmodel_auth_svc.UserDetails) (*responsemodel_auth_svc.SignupData, error) {
	var userData responsemodel_auth_svc.SignupData
	query := "INSERT INTO users (name, email, phone, password) VALUES($1, $2, $3, $4) RETURNING *"
	result := d.DB.Raw(query, userDetails.Name, userDetails.Email, userDetails.Phone, userDetails.Password).Scan(&userData)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userData, nil
}
