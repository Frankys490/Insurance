package redis_db

import (
	"Insurance/internal/model"
	"Insurance/pkg/custom_errors"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/valyala/fasthttp"
)

type RedisApiImpl struct {
	db11 *redis.Client
	db12 *redis.Client
}

func NewRedisApi(db11, db12 *redis.Client) *RedisApiImpl {
	return &RedisApiImpl{
		db11: db11,
		db12: db12,
	}
}

func (a *RedisApiImpl) CreateRedisEmployeeSingInDB11(req *model.CreateRedisEmployeeSingInDB11Req) *custom_errors.ErrHttp {
	err := a.db11.HSet(context.Background(), req.Key, map[string]interface{}{
		"password": req.Password,
		"active":   req.Active,
		"online":   req.Online,
		"salt":     req.Salt,
		"id":       req.ID,
		"ip":       req.IP,
	}).Err()
	if err != nil {
		return custom_errors.New(fasthttp.StatusInternalServerError, "post user data to redis: "+err.Error())
	}

	return nil
}

func (a *RedisApiImpl) GetRedisEmployeeSingInDB11(key string) (map[string]string, error) {
	values := a.db11.HGetAll(context.Background(), key).Val()
	if values == nil {
		return nil, fmt.Errorf("hgetall: %w", custom_errors.ErrNotFound)
	}

	a.db11.Del(context.Background(), key)

	return values, nil
}

func (a *RedisApiImpl) CreateRedisAuthSingInDB12(req *model.CreateRedisAuthSingInDB12Req) error {
	err := a.db12.HSet(context.Background(), req.Key, map[string]interface{}{
		"auth_code": req.AuthCode,
		"id":        req.ID,
		"ip":        req.IP,
	}).Err()
	if err != nil {
		return fmt.Errorf(fmt.Errorf("HSet: %w", err).Error()+": %w", custom_errors.ErrInternal)
	}

	return nil
}
