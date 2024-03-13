package models

type OTP struct {
	Otp string `bson:"otp" json:"otp"`
}

func NewOTP(otp string) *OTP {
	return &OTP{
		Otp: otp,
	}
}
