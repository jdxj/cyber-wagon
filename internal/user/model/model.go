package model

import (
	"context"
	"errors"
	"math/rand"

	"gorm.io/gorm"

	dao2 "github.com/jdxj/cyber-wagon/internal/user/dao"
	"github.com/jdxj/cyber-wagon/internal/user/proto"
)

func SignUp(ctx context.Context, req *proto.SignUpReq) (*proto.SignUpRsp, error) {
	user := &dao2.User{
		// todo
		Model:    gorm.Model{ID: uint(rand.Uint64())},
		Nickname: req.GetNickname(),
		Email:    req.GetEmail(),
		// todo
		Salt:     "",
		Password: req.GetPassword(),
	}
	err := dao2.DB.WithContext(ctx).
		Create(user).Error
	if err != nil {
		return nil, err
	}

	return &proto.SignUpRsp{UserId: uint64(user.ID)}, nil
}

func AuthBasic(ctx context.Context, req *proto.AuthBasicReq) (*proto.AuthBasicRsp, error) {
	user := &dao2.User{
		Email:    req.GetEmail(),
		Salt:     "",
		Password: req.GetPassword(),
	}
	err := dao2.DB.WithContext(ctx).
		Find(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.AuthBasicRsp{Allow: false}, nil
		}
		return nil, err
	}

	return &proto.AuthBasicRsp{Allow: true}, nil
}
