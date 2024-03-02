package domain_auth_svc

type Users struct {
	ID          uint `gorm:"unique;not null; primary key"`
	Name        string
	Email       string
	Phone       string
	Password    string
	ReferalCode string
	Status      string `gorm:"default:pending"`
}
