//TODO: NOTICE: this is usually placed in datamodels package of go-common lib!
package models

import (
	"fmt"
)

type Note struct {
	tableName struct{} `sql:"notes"`

	Id     int    `json:"noteId" pg:"id"`
	UserId int    `json:"userId" pg:"user_id"`
	Note   string `json:"note" pg:"note"`
}

func (t Note) String() string {
	return fmt.Sprintf("Note<%d>", t.Id)
}

func (t Note) Validate() error {

	// TODO: add here some expected validation rules

	return nil
}
