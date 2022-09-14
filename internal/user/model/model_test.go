package model

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jdxj/cyber-wagon/config"
	dao2 "github.com/jdxj/cyber-wagon/internal/user/dao"
	"github.com/jdxj/cyber-wagon/internal/user/proto"
)

func TestMain(m *testing.M) {
	c := config.DB{
		Host:   "192.168.50.200",
		Port:   3306,
		User:   "root",
		Pass:   "123456",
		DBName: "im",
	}
	dao2.Init(c)
	os.Exit(m.Run())
}

func TestSignUp(t *testing.T) {
	rsp, err := SignUp(context.Background(), &proto.SignUpReq{
		Email:    "example@gmail.com",
		Password: "123",
		Nickname: "abc",
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestAuthBasic(t *testing.T) {
	rsp, err := AuthBasic(context.Background(), &proto.AuthBasicReq{
		Email:    "example@gmail.com",
		Password: "123",
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}
