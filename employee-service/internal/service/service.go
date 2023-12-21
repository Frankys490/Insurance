package service

import (
	"Insurance/internal/api_db/redis_db"
	"Insurance/internal/api_db/reindexer_db"
	"Insurance/internal/model"
	"Insurance/pkg/custom_errors"
)

type AuthService interface {
	GenerateHashSingInService(req *model.GenerateHashHandlerReq) (*model.GenerateHashHandlerRes, *custom_errors.ErrHttp)
}

type Service struct {
	AuthService
}

func NewService(employeeApiDB *reindexer_db.EmployeeApiDB, redisApiDB *redis_db.RedisApiDB) *Service {
	return &Service{
		AuthService: NewAuthService(employeeApiDB.EmployeeApi, redisApiDB.RedisApi),
	}
}
