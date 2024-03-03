package requestmodel_auth_svc

type UserDetails struct {
	Id              string `json:"id"`
	Name            string `json:"name"           validate:"required"`
	Email           string `json:"email"          validate:"email"`
	Phone           string `json:"phone"          validate:"len=10"`
	Password        string `json:"password,omitempty"       validate:"min=4"`
	ConfirmPassword string `json:"confirmpassword,omitempty" validate:"eqfield=Password"`
}

type OtpVerification struct {
	Otp             string `json:"otp"   validate:"len=6"`
	TemporvaryToken string `validation:"required"`
}

type UserLogin struct {
	Phone    string `json:"phone"    validate:"len=10,number"`
	Password string `json:"password" validate:"required,min=4"`
}

type AdminLoginData struct {
    Email    string `json:"email"    validate:"email"`
    Password string `json:"password" validate:"min=4"`
}