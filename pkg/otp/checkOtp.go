package otp

import (
	"fmt"
	"shool/models"
	"shool/utils"
)

func CheckOTP(otp, hasedOtp *models.OTP) bool {

	if otp == nil {
		return false
	}

	fmt.Println(otp.Otp, hasedOtp.Otp)

	return utils.ComparePassword(otp.Otp, hasedOtp.Otp)

}
