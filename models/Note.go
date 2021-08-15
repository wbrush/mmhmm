//TODO: NOTICE: this is usually placed in datamodels package of go-common lib!
package models

import (
	"errors"
	"fmt"
)

// swagger:model NoteStruct
type Note struct {
	tableName struct{} `sql:"notes"`

	Id     int    `json:"noteId" pg:"id"`
	UserId int    `json:"-" pg:"user_id"`
	Author *User  `json:"author" pg:"-"`
	Note   string `json:"note" pg:"note"`
}

func (t Note) String() string {
	return fmt.Sprintf("Note<%d>", t.Id)
}

func (t Note) Validate() error {

	// TODO: add here some expected validation rules
	if t.Author == nil {
		return errors.New("Note does not have an Author!")
	}

	return nil
}
