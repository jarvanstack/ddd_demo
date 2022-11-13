package services

import (
	"context"
	"ddd_demo/application"
	"ddd_demo/domain"
	"ddd_demo/interfaces/rpc/protos/in/user"
)

var _ user.UserServer = &UserServerImpl{}

type UserServerImpl struct {
	UserApp *application.UserApp
	user.UnimplementedUserServer
}

func NewUserServer(userApp *application.UserApp) *UserServerImpl {
	return &UserServerImpl{
		UserApp: userApp,
	}
}

func (u *UserServerImpl) GetUser(ctx context.Context, req *user.GetUserReq) (*user.GetUserResp, error) {
	userID := req.GetId()

	if err := domain.ValidateUserID(userID); err != nil {
		return nil, err
	}

	userInfo, err := u.UserApp.Get(userID)
	if err != nil {
		return nil, err
	}

	return ToUserResp(userInfo), nil
}

func ToUserResp(u *domain.S2C_UserInfo) *user.GetUserResp {
	return &user.GetUserResp{
		User: &user.UserDTO{
			Id:       u.UserID,
			Username: u.Username,
		},
	}
}
