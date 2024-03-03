package responsemodel_auth_svc

type SignupData struct {
	ID          string `json:"userID"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Token       string `json:"token,omitempty"`
	RefrestToke string
}

type ValidatonError struct {
	Error string
}

type OtpValidation struct {
    Phone        string `json:"phone,omitempty"`
    Otp          string `json:"otp,omitempty"`
    Result       string `json:"result,omitempty"`
    Token        string `json:"token,omitempty"`
    AccessToken  string `json:"accesstoken,omitempty"`
    RefreshToken string `json:"refreshtoken,omitempty"`
}