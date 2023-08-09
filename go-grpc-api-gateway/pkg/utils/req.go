package utils

// input pagination
type Pagination struct {
	Offset uint `json:"offset"`
	Limit  uint `json:"limit"`
}

// input sign up user sign up details
type UsersSignUp struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	PhoneNum  string `json:"phonenum" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

// input admin sign up details

type AdminSignUp struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	PhoneNum  string `json:"Phonenum" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

// input verification of otp in case of forgot password
type OtpVerify struct {
	Otp         string `json:"otp" binding:"required"`
	OtpID       string `json:"otpid" binding:"required"`
	NewPassword string `json:"newpassword"`
}

// input verification of otp
type OtpSignUpVerify struct {
	Otp   string `json:"otp" binding:"required"`
	OtpID string `json:"otpid" binding:"required"`
}

// input login credentials

type LoginBody struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// input verification of otp at time of email or phone verification
type OtpLogin struct {
	Email    string `json:"email"`
	PhoneNum string `json:"phonenum"`
}
