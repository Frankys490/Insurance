package model

import (
	"Insurance/pkg/custom_errors"
)

type Response struct {
	Data    any    `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type GenerateHashHandlerReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IP       string
}

type GenerateHashHandlerRes struct {
	Username string `json:"username"`
	Hash     string `json:"hash"`
}

type SendEmailSingInHandlerReq struct {
	Username string `json:"username"`
	Hash     string `json:"hash"`
	IP       string
}

func (req *GenerateHashHandlerReq) Validate() *custom_errors.ErrHttp {
	if req.Username == "" || req.Password == "" {
		return custom_errors.ErrEmptyFields
	}

	return nil
}

func (req *SendEmailSingInHandlerReq) Validate() *custom_errors.ErrHttp {
	if req.Username == "" || req.Hash == "" {
		return custom_errors.ErrEmptyFields
	}

	return nil
}
