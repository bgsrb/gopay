package gopay

type IPaymentGateway interface {
	PayRequest(amount int,callback string) error
	VerifyRequest(amount int,callback string) error
	SettleRequest(amount int,callback string) error
	Finalize(amount int,callback string) error
	Redirect(amount int,callback string) error
}

// type Pay struct {
// }

// func NewPay(gateway IPaymentGateway) IPaymentGateway{
// 	return Mellat {
// 		terminalId : terminalId,
// 		userName : userName,
// 		userPassword : userPassword,
// 		client : webservice.NewIPaymentGateway("https://bpm.shaparak.ir/pgwchannel/services/pgw?wsdl",false, nil),
// 	}
// }