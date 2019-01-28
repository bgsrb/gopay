package mellat

//https://github.com/sinabakh/mellat-checkout/blob/master/src/mellat.js
//https://github.com/PoolPort/PoolPort/blob/master/src/Mellat/Mellat.php
//https://github.com/9px/MellatBank/blob/master/MellatBank.php
//https://github.com/Nevercom/php-ipg-ir

//https://github.com/iamtartan/epayment/blob/master/src/Epayment/Adapter/Mellat.php

//https://github.com/radoslav/soap

import (
	"strings"
	"time"

	ws "github.com/bgsrb/gopay/mellat/webservice"
)

var (
	wsdl   = "https://bpm.shaparak.ir/pgwchannel/services/pgw?wsdl"
	client *ws.IPaymentGateway
)

type IPaymentGateway interface {
	PayRequest(amount int, callback string) error
	VerifyRequest(amount int, callback string) error
	SettleRequest(amount int, callback string) error
}

type Mellat struct {
	terminalId   int64
	userName     string
	userPassword string
}

func New(terminalId int64, userName string, userPassword string) Mellat {
	return Mellat{
		terminalId:   terminalId,
		userName:     userName,
		userPassword: userPassword,
	}
}

// PayRequest returns ( refId,errorCode,error )
func (service Mellat) PayRequest(amount int64, orderId int64, callBackUrl string) (string, error) {
	req := &ws.BpPayRequest{
		TerminalId:     service.terminalId,
		UserName:       service.userName,
		UserPassword:   service.userPassword,
		OrderId:        orderId,
		Amount:         amount,
		LocalDate:      time.Now().Format("20091008"),
		LocalTime:      time.Now().Format("102003"),
		AdditionalData: "",
		CallBackUrl:    callBackUrl,
		PayerId:        0,
	}
	res, err := client.BpPayRequest(req)
	if err != nil {
		return "", err
	}
	ret := strings.Split(res.Return_, ",")
	if ret[0] == "0" {
		return ret[1], nil
	}
	return "", MellatError{code: ret[0]}
}

// VerifyRequest returns ( status, error )
func (service Mellat) VerifyRequest(orderId int64, saleOrderId int64, saleReferenceId int64) (bool, error) {
	req := &ws.BpVerifyRequest{
		TerminalId:      service.terminalId,
		UserName:        service.userName,
		UserPassword:    service.userPassword,
		OrderId:         orderId,
		SaleOrderId:     saleOrderId,
		SaleReferenceId: saleReferenceId,
	}
	res, err := client.BpVerifyRequest(req)
	if err != nil {
		return false, err
	}
	if res.Return_ == "0" {
		return true, nil
	}
	return false, nil
}

// SettleRequest returns ( status, error )
func (service Mellat) SettleRequest(orderId int64, saleOrderId int64, saleReferenceId int64) (bool, error) {
	req := &ws.BpSettleRequest{
		TerminalId:      service.terminalId,
		UserName:        service.userName,
		UserPassword:    service.userPassword,
		OrderId:         orderId,
		SaleOrderId:     saleOrderId,
		SaleReferenceId: saleReferenceId,
	}
	res, err := client.BpSettleRequest(req)
	if err != nil {
		return false, err
	}
	if res.Return_ == "0" || res.Return_ == "45" {
		return true, nil
	}
	return false, nil
}

func init() {
	client = ws.NewIPaymentGateway(wsdl, false, nil)
}
