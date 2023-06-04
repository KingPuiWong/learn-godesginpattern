package main

import (
	"fmt"
	"godesignpattern/behavior_patterns/template_method"
)

func main() {
	smsOTP := &template_method.Sms{}
	o := template_method.Otp{
		IOtp: smsOTP,
	}
	o.GenAndSendOTP(4)

	fmt.Println("")
	emailOTP := &template_method.Email{}
	o = template_method.Otp{IOtp: emailOTP}
	o.GenAndSendOTP(4)

}
