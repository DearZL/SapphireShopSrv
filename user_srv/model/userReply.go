package model

import "SapphireShop/SapphireShop_srv/user_srv/proto/srv"

func UserToResp(user User) *srv.UserInfoResponse {
	userInfoResp := &srv.UserInfoResponse{
		Id:        user.Id,
		UserId:    user.UserId,
		Email:     user.Email,
		UserName:  user.UserName,
		Birthday:  user.Birthday,
		Sex:       int32(user.Sex),
		Role:      int32(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return userInfoResp
}
