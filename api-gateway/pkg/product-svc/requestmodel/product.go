package requestmodel_product_svc_clind

type Category struct {
	Name string `json:"name" validate:"required,alpha"`
}

type Brand struct {
	Name string `json:"name" validate:"required,alpha"`
}