package service

import (
	"design-pattern/template-method/internal"
)

type Otp struct {
	IOtp internal.IOtp
}

func (o *Otp) GetAndSendOtp(otpLength int) error {
	otp := o.IOtp.GenRandomOTP(otpLength)
	o.IOtp.SaveOTPCache(otp)
	message := o.IOtp.GetMessage(otp)
	err := o.IOtp.SendNotification(message)
	if err != nil {
		return err
	}
	o.IOtp.PublishMetric()
	return nil
}
