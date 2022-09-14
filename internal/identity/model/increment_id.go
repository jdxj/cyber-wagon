package model

import (
	"context"
	"strconv"

	identity "github.com/jdxj/cyber-wagon/internal/identity/proto"
)

func IncrementID(ctx context.Context, req *identity.IncrementIDReq) (*identity.IncrementIDRsp, error) {
	hKey := "increment_id"
	key := strconv.FormatUint(req.GetNamespace(), 10)

	intCmd := redisClient.HIncrBy(ctx, hKey, key, req.GetQuantity())
	id, err := intCmd.Result()
	if err != nil {
		return nil, err
	}
	return &identity.IncrementIDRsp{MaxId: uint64(id)}, nil
}
