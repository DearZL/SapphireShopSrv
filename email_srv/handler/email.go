package handler

import (
	"SapphireShop/SapphireShop_srv/email_srv/global"
	"SapphireShop/SapphireShop_srv/email_srv/proto/srv"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/gomail.v2"
	"time"
)

type EmailServer struct {
	srv.UnimplementedEmailSrvServer
}

func (s *EmailServer) SendCode(ctx context.Context, req *srv.Email) (*emptypb.Empty, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(
		global.Config.GetString("emailConfig.username"),
		"Sapphire商城"))
	m.SetHeader("To", req.Email)
	m.SetHeader("Subject", req.Subject)
	m.SetBody("text/html", req.Msg)
	d := gomail.NewDialer(
		global.Config.GetString("emailConfig.server"),
		global.Config.GetInt("emailConfig.port"),
		global.Config.GetString("emailConfig.username"),
		global.Config.GetString("emailConfig.key"))
	err := d.DialAndSend(m)
	if err != nil {
		global.Logger.Info("邮件发送失败")
		global.Logger.Info(err.Error())
		return nil, err
	}
	if req.Subject == "验证码" {
		global.Redis.Set(req.Email, req.Code, time.Duration(req.Expire))
	}
	global.Logger.Info("邮件发送成功")
	return &emptypb.Empty{}, nil
}
