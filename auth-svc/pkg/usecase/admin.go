package usecase_auth_svc

import (
	"fmt"

	config_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/config"
	requestmodel_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/model/requestmodel"
	responsemodel_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/model/responsemodel"
	interfaceRepository_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/repository/interface"
	service_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/services"
	interfaceUseCase_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/usecase/interface"
	helper_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/utils"
	"google.golang.org/protobuf/types/known/emptypb"
)

type adminUsecase struct {
	repo             interfaceRepository_auth_svc.IAdminRepository
	tokenSecurityKey config_auth_svc.Token
}

func NewAdminUseCase(adminRepository interfaceRepository_auth_svc.IAdminRepository, key *config_auth_svc.Token) interfaceUseCase_auth_svc.IAdminUseCase {
	return &adminUsecase{repo: adminRepository,
		tokenSecurityKey: *key}
}

func (r *adminUsecase) AdminLogin(adminData *requestmodel_auth_svc.AdminLoginData) (*responsemodel_auth_svc.AdminLoginRes, error) {
	var adminLoginRes responsemodel_auth_svc.AdminLoginRes

	HashedPassword, err := r.repo.GetPassword(adminData.Email)
	if err != nil {
		return nil, err
	}

	err = helper_auth_svc.CompairPassword(HashedPassword, adminData.Password)
	if err != nil {
		return nil, err
	}

	token, err := service_auth_svc.GenerateRefreshToken(r.tokenSecurityKey.AdminSecurityKey)
	if err != nil {
		return nil, err
	}

	adminLoginRes.Token = token
	return &adminLoginRes, nil
}

func (u *adminUsecase) VerifyAdminToken(token string) (*emptypb.Empty, error) {
	fmt.Println("admin middlewire auth")
	err := service_auth_svc.VerifyRefreshToken(token, u.tokenSecurityKey.AdminSecurityKey)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
