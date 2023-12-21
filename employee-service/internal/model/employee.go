package model

import "github.com/restream/reindexer/v3"

type EmployeeItem struct {
	ID                    int64           `json:"id" reindex:"id,hash,pk"`
	Name                  string          `json:"name" reindex:"name,hash"`
	Surname               string          `json:"surname" reindex:"surname,hash"`
	Patronymic            string          `json:"patronymic" reindex:"patronymic,hash"`
	Username              string          `json:"username" reindex:"username,hash"`
	Password              string          `json:"password" reindex:"password,hash"`
	LastIP                string          `json:"last_ip" reindex:"last_ip,hash"`
	LastEntryPlace        reindexer.Point `json:"last_entry_place" reindex:"last_entry_place,rtree"`
	ActiveStatus          int32           `json:"active_status" reindex:"active_status,hash"`
	NumberOfWrongEntryTry int64           `json:"number_of_wrong_entry_try" reindex:"number_of_wrong_entry_try,hash"`
	EmployeeRole          string          `json:"employee_role" reindex:"employee_role,hash"`
	RuleExternalEntry     int32           `json:"rule_external_entry" reindex:"rule_external_entry,hash"`
	LastEntryDate         int64           `json:"last_entry_date" reindex:"last_entry_date,tree"`
}
