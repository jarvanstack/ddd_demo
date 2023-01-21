package test

import (
	"context"
	pb_user "ddd_demo/internal/servers/rpc/protos/in/user"
	"fmt"
	"testing"

	"google.golang.org/grpc"
)

// grpc 客户端调用测试
func Test_Rpc_UserInfo(t *testing.T) {
	conn, err := grpc.Dial(":8889", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}

	resp, err := pb_user.NewUserClient(conn).UserInfo(context.Background(), &pb_user.G2S_UserInfo{
		Id: "1",
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("resp: %v\n", resp)
}
