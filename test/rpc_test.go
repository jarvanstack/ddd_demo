package main

import (
	"context"
	"ddd_demo/internal/interfaces/rpc/protos/in/user"
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

	resp, err := user.NewUserClient(conn).GetUser(context.Background(), &user.GetUserReq{
		Id: "1",
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("resp: %v\n", resp)
}
