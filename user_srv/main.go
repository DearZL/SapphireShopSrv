package main

import (
	"SapphireShop/SapphireShop_srv/user_srv/common/config"
	"SapphireShop/SapphireShop_srv/user_srv/global"
	"SapphireShop/SapphireShop_srv/user_srv/handler"
	"SapphireShop/SapphireShop_srv/user_srv/proto/srv"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func init() {
	global.InitConfig()
	global.InitLogConfig()
	global.InitDB()
	global.InitRedis()
}

func main() {
	//结束时刷新log缓冲区到
	defer config.SyncLog(global.Logger)
	ip := global.Config.GetString("userServer.host")
	port := global.Config.GetInt("userServer.port")
	server := grpc.NewServer()
	srv.RegisterUserSrvServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		panic(err)
	}
	global.Logger.Info("UserSrv start server on " + fmt.Sprintf("%s:%d", ip, port))
	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
