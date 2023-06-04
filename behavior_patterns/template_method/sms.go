package template_method

import "fmt"

type Sms struct {
	Otp
}

func (s Sms) GenRandomOPT(len int) string {
	randomOpt := "1234"
	fmt.Printf("SMS: generating random opt %s\n", randomOpt)
	return randomOpt
}

func (s Sms) SaveOTPCache(otp string) {
	fmt.Printf("SMS: saving opt: %s to cache\n", otp)
}

func (s Sms) GetMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s Sms) SendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}
