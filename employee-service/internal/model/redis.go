package model

type CreateRedisEmployeeSingInDB11Req struct {
	Key      string
	Password bool
	Active   bool
	Online   bool
	Salt     string
	ID       int64
	IP       string
}

type CreateRedisAuthSingInDB12Req struct {
	Key      string
	AuthCode int
	ID       int64
	IP       string
}
