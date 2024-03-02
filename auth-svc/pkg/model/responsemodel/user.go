package responsemodel_auth_svc

type SignupData struct {
	ID          string `json:"userID"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Token       string `json:"token,omitempty"`
	RefrestToke string
}
