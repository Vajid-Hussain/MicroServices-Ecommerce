package server_auth_svc

import (
	"context"
	"fmt"

	requestmodel_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/model/requestmodel"
	"github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/pb"
	interfaceUseCase_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/usecase/interface"
	helper_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/utils"
)

type UserService struct {
	useCase interfaceUseCase_auth_svc.IuserUseCase
	pb.UnimplementedAuthServiceServer
}

func NewUserService(usecase interfaceUseCase_auth_svc.IuserUseCase) *UserService {
	return &UserService{useCase: usecase}
}

func (u *UserService) Sighup(ctx context.Context, req *pb.SignupRequest) (*pb.SignupResponse, error) {
	var userDetails requestmodel_auth_svc.UserDetails
	userDetails.ConfirmPassword = req.ConfirmPassword
	userDetails.Password = req.Password
	userDetails.Email = req.Email
	userDetails.Name = req.Name
	userDetails.Phone = req.Phone

	validationErr := helper_auth_svc.EasyValidataion(userDetails)

	var errorString string
	if len(validationErr) > 0 {
		for _, val := range validationErr {
			errorString += fmt.Errorf(val.Error).Error()
		}
	}
	if len(validationErr) > 0 {
		return nil, fmt.Errorf(errorString)
	}

	result, err := u.useCase.UserSignup(&userDetails)
	if err != nil {
		return nil, err
	}

	return &pb.SignupResponse{ID: result.ID,
		Name:  result.Name,
		Email: result.Email,
		Phone: result.Phone,
		Token: result.Token,
	}, nil
}

func (u *UserService) OtpVerifiction(ctx context.Context, req *pb.OtpRequest) (*pb.OtpResponse, error) {
	var otp requestmodel_auth_svc.OtpVerification
	otp.Otp = req.Otp
	otp.TemporvaryToken = req.TemperverToken
	fmt.Println("----", otp)

	validationErr := helper_auth_svc.EasyValidataion(otp)

	var errorString string
	if len(validationErr) > 0 {
		for _, val := range validationErr {
			errorString += fmt.Errorf(val.Error).Error()
		}
	}
	if len(validationErr) > 0 {
		return nil, fmt.Errorf(errorString)
	}

	result, err := u.useCase.VerifyOtp(otp, otp.TemporvaryToken)
	if err != nil {
		return nil, err
	}

	return &pb.OtpResponse{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
	}, nil
}
