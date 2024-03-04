package di_product_svc

import (
	config_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/config"
	db_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/db"
	repository_poduct_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/repository"
	server_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/server"
	usecase_prodcut_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/usecase"
)

func InitializeServer(config *config_product_svc.Config) (*server_product_svc.ProductHandler, error) {
	db, err := db_product_svc.InitDB(config)
	if err != nil {
		return nil, err
	}

	categoryRepository := repository_poduct_svc.NewCategoryRepository(db)
	categoryUseCase := usecase_prodcut_svc.NewCategoryUseCase(categoryRepository)
	server := server_product_svc.NewProductHandler(categoryUseCase)

	return server, nil
}
