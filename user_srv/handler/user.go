package handler

import (
	"SapphireShop/SapphireShop_srv/user_srv/global"
	"SapphireShop/SapphireShop_srv/user_srv/model"
	"SapphireShop/SapphireShop_srv/user_srv/proto/srv"
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/oklog/ulid/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type UserServer struct {
	srv.UnimplementedUserSrvServer
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (s *UserServer) GetUserList(ctx context.Context, req *srv.PageInfo) (*srv.UserListResponse, error) {
	//获取用户列表
	var users []model.User
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "无用户存在")
	}
	resp := &srv.UserListResponse{}
	resp.Total = result.RowsAffected
	global.DB.Scopes(Paginate(int(req.PageNo), int(req.PageSize))).Find(&users)
	for _, user := range users {
		resp.Data = append(resp.Data, model.UserToResp(user))
	}
	return resp, nil
}

func (s *UserServer) GetUserByMobile(ctx context.Context, req *srv.EmailRequest) (*srv.UserInfoResponse, error) {
	//通过Email查询用户
	var user model.User
	result := global.DB.Where(&model.User{Email: req.Email}).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	userResp := model.UserToResp(user)
	return userResp, nil
}

func (s *UserServer) GetUserById(ctx context.Context, req *srv.IdRequest) (*srv.UserInfoResponse, error) {
	//通过Id查询用户
	var user model.User
	result := global.DB.First(&user, req.UserId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	userResp := model.UserToResp(user)
	return userResp, nil
}

func (s *UserServer) CreateUser(ctx context.Context, req *srv.CreateUserInfo) (*srv.UserInfoResponse, error) {
	//比对验证码
	if global.Redis.Get(req.Email).Val() == "" {
		return nil, status.Errorf(codes.Aborted, "验证码已失效")
	}
	if global.Redis.Get(req.Email).Val() != req.Code {
		return nil, status.Errorf(codes.Aborted, "验证码错误")
	}
	//删除验证码
	global.Redis.Del(req.Email)
	//新建用户
	var user model.User
	result := global.DB.Where(&model.User{Email: req.Email}).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected != 0 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}
	user.UserId = ulid.Make().String()
	user.Email = req.Email
	user.UserName = req.UserName
	h := md5.Sum([]byte(req.Password))
	user.Password = hex.EncodeToString(h[:])
	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return model.UserToResp(user), nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *srv.UpdateUserInfo) (*emptypb.Empty, error) {
	//更新用户
	var user model.User
	result := global.DB.First(&user, req.UserId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	user.UserName = req.UserName
	user.Birthday = req.Birthday
	user.Sex = uint8(req.Sex)
	result = global.DB.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &emptypb.Empty{}, nil
}
