package template_method

type IOtp interface {
	GenRandomOPT(int) string
	SaveOTPCache(string)
	GetMessage(string) string
	SendNotification(string) error
}

type Otp struct {
	IOtp IOtp
}

func (o *Otp) GenAndSendOTP(optLength int) error {
	otp := o.IOtp.GenRandomOPT(optLength)
	o.IOtp.SaveOTPCache(otp)
	message := o.IOtp.GetMessage(otp)
	err := o.IOtp.SendNotification(message)
	if err != nil {
		return err
	}
	return nil
}
