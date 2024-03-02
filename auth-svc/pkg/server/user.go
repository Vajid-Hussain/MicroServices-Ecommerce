package server_auth_svc

import (
	"context"

	requestmodel_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/model/requestmodel"
	interfaceUseCase_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/usecase/interface"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/pb"
)

type UserService struct {
	useCase interfaceUseCase_auth_svc.IuserUseCase
}

func NewUserService(usecase interfaceUseCase_auth_svc.IuserUseCase) *UserService {
	return &UserService{useCase: usecase}
}

func (u *UserService) Sighup(ctx context.Context, req *pb.SignupRequest) (pb.SignupResponse, error) {
	var userDetails requestmodel_auth_svc.UserDetails
	userDetails.ConfirmPassword = req.ConfirmPassword
	userDetails.Password = req.Password
	userDetails.Email = req.Email
	userDetails.Name = req.Name
	userDetails.Phone = req.Phone

	result, err := u.useCase.UserSignup(&userDetails)
	return pb.SignupResponse{ID: result.ID,
		Name:  result.Name,
		Email: result.Email,
		Phone: result.Phone,
		Token: result.Token,
	}, err
}
