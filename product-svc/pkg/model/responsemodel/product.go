package responsemodel_product_svc


type CategoryDetails struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type BrandDetails struct {
	ID   string `json:"id" validate:"required,number"`
	Name string `json:"name" validate:"required"`
}

type ValidatonError struct {
    Error string
}