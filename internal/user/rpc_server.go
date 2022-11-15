package user

import (
	"context"
	"ddd_demo/internal/servers/rpc/protos/in/user"
	pb "ddd_demo/internal/servers/rpc/protos/in/user"
	"ddd_demo/internal/user/model"
)

var _ pb.UserServer = &UserRpcServerImpl{}

type UserRpcServerImpl struct {
	UserApp UserAppInterface
	pb.UnimplementedUserServer
}

func NewUserServer(userApp UserAppInterface) *UserRpcServerImpl {
	return &UserRpcServerImpl{
		UserApp: userApp,
	}
}

func (u *UserRpcServerImpl) GetUser(ctx context.Context, req *user.G2S_UserInfo) (*user.S2G_UserInfo, error) {
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

func ToUserResp(u *model.S2C_UserInfo) *user.S2G_UserInfo {
	return &user.S2G_UserInfo{
		User: &user.UserDTO{
			Id:       u.UserID,
			Username: u.Username,
		},
	}
}
