package redis_db

import (
	"Insurance/internal/model"
	"Insurance/pkg/custom_errors"
	"github.com/redis/go-redis/v9"
)

type RedisApi interface {
	CreateRedisEmployeeSingInDB11(req *model.CreateRedisEmployeeSingInDB11Req) *custom_errors.ErrHttp
	GetRedisEmployeeSingInDB11(key string) (map[string]string, error)
	CreateRedisAuthSingInDB12(req *model.CreateRedisAuthSingInDB12Req) error
}

type RedisApiDB struct {
	RedisApi
}

func NewRedisApiDB(db11, db12 *redis.Client) *RedisApiDB {
	return &RedisApiDB{
		RedisApi: NewRedisApi(db11, db12),
	}
}
