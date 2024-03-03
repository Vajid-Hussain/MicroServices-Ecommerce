package di_auth_svc

import (
	"log"

	config_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/config"
	db_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/db"
	repository_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/repository"
	server_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/server"
	service_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/services"
	usecase_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/usecase"
)

func InitializeServer(config config_auth_svc.Config) *server_auth_svc.UserService {
	db, err := db_auth_svc.InitDB(&config.General)
	if err != nil {
		log.Fatal(err, " fatal because db connection error")
	}

	service_auth_svc.OtpService(config.Otp)

	userRepository := repository_auth_svc.NewUserRepository(db)
	userUseCase := usecase_auth_svc.NesUserUseCase(userRepository, config.Token)

	adminRepository := repository_auth_svc.NewAdminRepository(db)
	adminUseCase := usecase_auth_svc.NewAdminUseCase(adminRepository, &config.Token)

	userServer := server_auth_svc.NewUserService(userUseCase, adminUseCase)

	return userServer

}
