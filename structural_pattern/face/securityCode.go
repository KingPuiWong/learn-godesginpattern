package face

import "fmt"

type SecurityCode struct {
	Code int
}

func NewSecurityCode(code int) *SecurityCode {
	return &SecurityCode{code}
}

func (s *SecurityCode) checkCode(incomingCode int) error {
	if s.Code != incomingCode {
		return fmt.Errorf("security Code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}
