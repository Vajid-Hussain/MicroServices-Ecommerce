package interfaceUseCase_auth_svc

import (
	requestmodel_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/model/requestmodel"
	responsemodel_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/model/responsemodel"
)

type IAdminUseCase interface {
	AdminLogin( *requestmodel_auth_svc.AdminLoginData) (*responsemodel_auth_svc.AdminLoginRes, error)
}

