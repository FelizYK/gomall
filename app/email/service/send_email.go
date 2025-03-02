package service

import (
	rpcemail "github.com/FelizYK/gomall/rpc/email"
	"github.com/kr/pretty"
)

func SendEmail(req *rpcemail.SendEmailReq) (err error) {
	pretty.Printf("SendEmail: %v\n", req)
	return
}
