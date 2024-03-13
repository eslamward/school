package otp

import (
	"math/rand"
	"shool/models"
	"strings"
)

func GenerateOTP() *models.OTP {
	var listnum []string
	for i := 0; i < 6; i++ {
		randInt := rand.Intn(len(charchter))
		char := string(charchter[randInt])
		listnum = append(listnum, char)
	}
	otpString := strings.Join(listnum, "")
	return models.NewOTP(otpString)
}
