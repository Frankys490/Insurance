package reindexer_db

import (
	"Insurance/internal/model"
	"Insurance/pkg/custom_errors"
	"github.com/restream/reindexer/v3"
)

type EmployeeApi interface {
	GetEmployeeByUsernameDB(username string) (*model.EmployeeItem, *custom_errors.ErrHttp)
}

type EmployeeApiDB struct {
	EmployeeApi
}

func NewEmployeeApiDB(db *reindexer.Reindexer) *EmployeeApiDB {
	return &EmployeeApiDB{
		EmployeeApi: NewEmployeeApi(db),
	}
}
