package usecase_auth_svc

import (
	"errors"

	config_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/config"
	requestmodel_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/model/requestmodel"
	responsemodel_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/model/responsemodel"
	interfaceRepository "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/repository/interface"
	service_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/services"
	interfaceUseCase_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/usecase/interface"
	helper_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/utils"
)

type userUseCase struct {
	repo  interfaceRepository.IuserRepository
	token config_auth_svc.Token
}

func NesUserUseCase(repo interfaceRepository.IuserRepository, tokenCredential config_auth_svc.Token) interfaceUseCase_auth_svc.IuserUseCase {
	return &userUseCase{
		repo:  repo,
		token: tokenCredential,
	}
}

func (u *userUseCase) UserSignup(userData *requestmodel_auth_svc.UserDetails) (*responsemodel_auth_svc.SignupData, error) {

	var resSignup *responsemodel_auth_svc.SignupData

	if isExist := u.repo.IsUserExist(userData.Phone); isExist >= 1 {
		return nil, errors.New("user is exist try again , with another phone number")
	} else {
		service_auth_svc.TwilioSetup()
		_, err := service_auth_svc.SendOtp(userData.Phone)
		if err != nil {
			return nil, errors.New("error at attempt for creating new OTP")
		}

		userData.Password = helper_auth_svc.HashPassword(userData.Password)

		resSignup, err = u.repo.CreateUser(userData)
		if err != nil {
			return nil, err
		}

		token, err := service_auth_svc.TemperveryTokenForOtpVerification(u.token.TemperveryKey, userData.Phone)
		if err != nil {
			return nil, errors.New("creating token is error")
		}
		resSignup.Token = token

	}

	return resSignup, nil
}
