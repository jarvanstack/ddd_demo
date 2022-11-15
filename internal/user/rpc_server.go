package user

import (
	"context"
	"ddd_demo/internal/servers/rpc/protos/in/user"
	pb "ddd_demo/internal/servers/rpc/protos/in/user"
	"ddd_demo/internal/user/model"
)

var _ pb.UserServer = &UserRpcServerImpl{}

type UserRpcServerImpl struct {
	UserApp *UserApp
	pb.UnimplementedUserServer
}

func NewUserServer(userApp *UserApp) *UserRpcServerImpl {
	return &UserRpcServerImpl{
		UserApp: userApp,
	}
}

func (u *UserRpcServerImpl) GetUser(ctx context.Context, req *user.GetUserReq) (*user.GetUserResp, error) {
	userID, err := model.NewUserID(req.Id)
	if err != nil {
		return nil, err
	}

	userInfo, err := u.UserApp.Get(userID)
	if err != nil {
		return nil, err
	}

	return ToUserResp(userInfo), nil
}

func ToUserResp(u *model.S2C_UserInfo) *user.GetUserResp {
	return &user.GetUserResp{
		User: &user.UserDTO{
			Id:       u.UserID,
			Username: u.Username,
		},
	}
}
