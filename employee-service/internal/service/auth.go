package service

import (
	"Insurance/internal/api_db/reindexer_db"
	"Insurance/internal/model"
	"Insurance/pkg/custom_errors"
	"crypto/rand"
	"encoding/hex"
	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

type AuthServiceImpl struct {
	employeeApi reindexer_db.EmployeeApi
	//redisApi    redis_db.RedisApi
}

func NewAuthService(employeeApi reindexer_db.EmployeeApi /*, redisApi redis_db.RedisApi*/) *AuthServiceImpl {
	return &AuthServiceImpl{
		employeeApi: employeeApi,
		//redisApi:    redisApi,
	}
}

func (s *AuthServiceImpl) GenerateHashSingInService(req *model.GenerateHashHandlerReq) (*model.GenerateHashHandlerRes, *custom_errors.ErrHttp) {
	salt := make([]byte, 256)

	employee, errCustom := s.employeeApi.GetEmployeeByUsernameDB(req.Username)
	if errCustom != nil {
		return nil, custom_errors.New(errCustom.Code, "get employee by username DB: "+errCustom.Message)
	}

	//var trialPass, trialActive, trialOnline bool

	err := bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(req.Password))
	/*if err == nil {
		trialPass = true
	}

	if employee.ActiveStatus == 1 {
		trialActive = true
	}*/

	_, err = rand.Read(salt)
	if err != nil {
		return nil, custom_errors.New(fasthttp.StatusInternalServerError, "salt generating: "+err.Error())
	}

	hashVer := sha3.Sum512([]byte(employee.Username + employee.Password + hex.EncodeToString(salt)))

	loss := &model.GenerateHashHandlerRes{
		Username: employee.Username,
		Hash:     hex.EncodeToString(hashVer[:]),
	}

	/*reqRedis := &model.CreateRedisEmployeeSingInDB11Req{
		Key:      employee.Username + "__" + hex.EncodeToString(hashVer[:]),
		Active:   trialActive,
		Password: trialPass,
		Online:   trialOnline,
		Salt:     hex.EncodeToString(salt),
		ID:       employee.ID,
		IP:       req.IP,
	}

	errCustom = s.redisApi.CreateRedisEmployeeSingInDB11(reqRedis)
	if errCustom != nil {
		return nil, custom_errors.New(errCustom.Code, "create redis employee sing in DB11: "+errCustom.Message)
	}*/

	return loss, nil
}
