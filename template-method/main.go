package main

import (
	"design-pattern/template-method/internal/service"
)

func main() {
	sms := &service.Sms{}
	otp := service.Otp{IOtp: sms}
	err := otp.GetAndSendOtp(6)
	if err != nil {
		return
	}

	email := &service.Email{}
	otp = service.Otp{IOtp: email}
	err = otp.GetAndSendOtp(5)
	if err != nil {
		return
	}

}
