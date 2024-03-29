package interfaceUseCase_auth_svc

import (
	requestmodel_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/model/requestmodel"
	responsemodel_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/model/responsemodel"
)

type IuserUseCase interface {
	UserSignup(*requestmodel_auth_svc.UserDetails) (*responsemodel_auth_svc.SignupData, error)
	VerifyOtp( requestmodel_auth_svc.OtpVerification,  string) (responsemodel_auth_svc.OtpValidation, error) 
	UserLogin( requestmodel_auth_svc.UserLogin) (responsemodel_auth_svc.UserLogin, error) 
	VerifyUserToken(string, string)(string,error)
}
