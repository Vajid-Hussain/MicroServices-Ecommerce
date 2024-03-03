package interfaceRepository_auth_svc

import (
	requestmodel_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/model/requestmodel"
	"github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/model/responsemodel"
)

type IuserRepository interface {
	CreateUser(*requestmodel_auth_svc.UserDetails) (*responsemodel_auth_svc.SignupData, error)
	IsUserExist(string) int
	FetchUserID(string) (string, error)
	ChangeUserStatusActive(string) error
	FetchPasswordUsingPhone( string) (string, error)

}
