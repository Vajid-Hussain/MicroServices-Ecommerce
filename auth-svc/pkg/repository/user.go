package repository_auth_svc

import (
	"database/sql"
	"errors"
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
	fmt.Println("--callded user adde.dc")
	var userData responsemodel_auth_svc.SignupData
	query := "INSERT INTO users (name, email, phone, password) VALUES($1, $2, $3, $4) RETURNING *"
	result := d.DB.Raw(query, userDetails.Name, userDetails.Email, userDetails.Phone, userDetails.Password).Scan(&userData)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userData, nil
}

func (d *userRepository) FetchUserID(phone string) (string, error) {
	var userID string

	query := "SELECT id FROM users WHERE phone=? AND status='active'"
	data := d.DB.Raw(query, phone).Row()

	if err := data.Scan(&userID); err != nil {
		return "", errors.New("fetching user id cause error")
	}
	return userID, nil
}

func (d *userRepository) ChangeUserStatusActive(phone string) error {
	fmt.Println(phone)
	query := "UPDATE users SET status = 'active' WHERE phone = ?"
	result := d.DB.Exec(query, phone)
	if result.Error != nil {

		return errors.New("no user Exist , phone number is wrong")
	} else {
		return nil
	}
}

func (d *userRepository) FetchPasswordUsingPhone(phone string) (string, error) {
	var password string

	query := "SELECT password FROM users WHERE phone=? AND status='active'"
	row := d.DB.Raw(query, phone).Row()
	fmt.Println("--------", row)

	if row == nil {
		return "", errors.New("no user exist or you are blocked by admin")
	}

	err := row.Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("user does not exist or user get blocked")
		}
		return "", fmt.Errorf("error scanning row: %s", err)
	}
	return password, nil
}
