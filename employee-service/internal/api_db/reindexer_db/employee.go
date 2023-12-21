package reindexer_db

import (
	"Insurance/internal/model"
	"Insurance/pkg/custom_errors"
	"github.com/restream/reindexer/v3"
)

type EmployeeApiImpl struct {
	db *reindexer.Reindexer
}

func NewEmployeeApi(db *reindexer.Reindexer) *EmployeeApiImpl {
	return &EmployeeApiImpl{
		db: db,
	}
}

func (a *EmployeeApiImpl) GetEmployeeByUsernameDB(username string) (*model.EmployeeItem, *custom_errors.ErrHttp) {
	elem, found := a.db.Query("employees").WhereString("username", reindexer.EQ, username).Get()

	if !found {
		return nil, custom_errors.ErrWrongUserInputData
	}

	employee := elem.(*model.EmployeeItem)
	return employee, nil
}
