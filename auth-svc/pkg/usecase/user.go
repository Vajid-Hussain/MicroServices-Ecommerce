package usecase_auth_svc

import (
	"errors"
	"fmt"

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
	fmt.Println("readh usecase", u.token)
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

func (u *userUseCase) VerifyOtp(otpConstrain requestmodel_auth_svc.OtpVerification, token string) (responsemodel_auth_svc.OtpValidation, error) {

	var otpResponse responsemodel_auth_svc.OtpValidation
	phone, err := service_auth_svc.FetchPhoneFromToken(token, u.token.TemperveryKey)
	if err != nil {
		otpResponse.Token = "invalid token"
		return otpResponse, errors.New("error at token extraction, cause by invalid token")
	}
	service_auth_svc.TwilioSetup()

	if err := service_auth_svc.VerifyOtp(phone, otpConstrain.Otp); err != nil {
		otpResponse.Otp = "otp verification failed"
		return otpResponse, errors.New("otp is not match with your number, try egain")
	}

	if err := u.repo.ChangeUserStatusActive(phone); err != nil {
		otpResponse.Phone = "no user exist with phone number , verify is phone number is it correct "
		return otpResponse, errors.New("no user exist ")
	}

	userID, err := u.repo.FetchUserID(phone)
	if err != nil {
		otpResponse.Result = "error at db attempt to featch userID"
		return otpResponse, errors.New("error cause by fetching user id")
	}

	accessToken, err := service_auth_svc.GenerateAcessToken(u.token.UserSecurityKey, userID)
	if err != nil {
		otpResponse.Result = "creating token not done succesfully"
		return otpResponse, errors.New("token creation cause error")
	}

	refreshToken, err := service_auth_svc.GenerateRefreshToken(u.token.UserSecurityKey)
	if err != nil {
		otpResponse.Result = "creating token not done succesfully"
		return otpResponse, errors.New("token creation cause error")
	}

	otpResponse.AccessToken = accessToken
	otpResponse.RefreshToken = refreshToken
	otpResponse.Result = "success"

	return otpResponse, nil
}

func (u *userUseCase) UserLogin(loginCredential requestmodel_auth_svc.UserLogin) (responsemodel_auth_svc.UserLogin, error) {
	var resUserLogin responsemodel_auth_svc.UserLogin

	password, err := u.repo.FetchPasswordUsingPhone(loginCredential.Phone)
	if err != nil {
		resUserLogin.Error = err.Error()
		return resUserLogin, err
	}

	err = helper_auth_svc.CompairPassword(password, loginCredential.Password)
	if err != nil {
		resUserLogin.Error = err.Error()
		return resUserLogin, err
	}

	userID, err := u.repo.FetchUserID(loginCredential.Phone)
	if err != nil {
		resUserLogin.Error = err.Error()
		return resUserLogin, err
	}

	accessToken, err := service_auth_svc.GenerateAcessToken(u.token.UserSecurityKey, userID)
	if err != nil {
		resUserLogin.Error = err.Error()
		return resUserLogin, err
	}

	refreshToken, err := service_auth_svc.GenerateRefreshToken(u.token.UserSecurityKey)
	if err != nil {
		resUserLogin.Error = err.Error()
	}

	resUserLogin.AccessToken = accessToken
	resUserLogin.RefreshToken = refreshToken

	return resUserLogin, nil
}

func (u *userUseCase) VerifyUserToken(accessToken, refreshToken string) (string, error) {
	fmt.Println("user middlewiere")
	id, err := service_auth_svc.VerifyAcessToken(accessToken, u.token.UserSecurityKey)
	if err != nil {
		fmt.Println("error at accestoken", err)
	}

	err = service_auth_svc.VerifyRefreshToken(refreshToken, u.token.UserSecurityKey)
	if err != nil {
		return "", err
	}
	return id, nil
}


