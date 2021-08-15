package models

import (
	"fmt"
)

// swagger:model UserStruct
type User struct {
	tableName struct{} `sql:"users"`

	Id        int    `json:"userId" pg:"id"`
	FirstName string `json:"firstName" pg:"first_name"`
	LastName  string `json:"lastName" pg:"last_name"`
}

func (t User) String() string {
	return fmt.Sprintf("User<%d, %s %s>", t.Id, t.FirstName, t.LastName)
}

func (t User) Validate() error {

	// TODO: add here some expected validation rules

	return nil
}
