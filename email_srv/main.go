package main

import (
	"SapphireShop/SapphireShop_srv/email_srv/common/config"
	"SapphireShop/SapphireShop_srv/email_srv/global"
	"SapphireShop/SapphireShop_srv/email_srv/handler"
	"SapphireShop/SapphireShop_srv/email_srv/proto/srv"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func init() {
	global.InitConfig()
	global.InitLogConfig()
	global.InitRedis()
}

func main() {
	//结束时刷新log缓冲区到
	defer config.SyncLog(global.Logger)
	ip := global.Config.GetString("emailServer.host")
	port := global.Config.GetInt("emailServer.port")
	server := grpc.NewServer()
	srv.RegisterEmailSrvServer(server, &handler.EmailServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		panic(err)
	}
	global.Logger.Info("EmailSrv start server on " + fmt.Sprintf("%s:%d", ip, port))
	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
