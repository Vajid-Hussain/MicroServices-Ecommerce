package di_auth_svc

import (
	"log"

	config_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/config"
	db_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/db"
	repository_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/repository"
	usecase_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/usecase"
)

func InitializeServer(config config_auth_svc.Config) {
	db, err := db_auth_svc.InitDB(&config.DB)
	if err != nil {
		log.Fatal(err)
	}

	userRepository:= repository_auth_svc.NewUserRepository(db)
	userUseCase:= usecase_auth_svc.NesUserUseCase(userRepository, config.Token)

}
