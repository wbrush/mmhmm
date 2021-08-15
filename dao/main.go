package dao

import (
	"net/url"

	"github.com/wbrush/go-common/db"
	"github.com/wbrush/mmhmm/models"
)

type (
	Notes interface {
		CreateNote(note *models.Note) (isDuplicate bool, err error)
		GetNoteById(id int) (note *models.Note, isFound bool, err error)
		ListNotes(filters url.Values) (list []models.Note, err error)
		UpdateNoteById(note *models.Note) (err error)
		DeleteNoteById(id int) (isFound bool, err error)
	}

	Users interface {
		CreateUser(user *models.User) (isDuplicate bool, err error)
		GetUserById(id int) (user *models.User, isFound bool, err error)
		ListUsers(filters url.Values) (list []models.User, err error)
		UpdateUserById(user *models.User) (err error)
		DeleteUserById(id int) (isFound bool, err error)
	}

	DataAccessObject interface {
		db.BaseDataAccessObject //this is need only if you plain to use transactions or some base additional features
		Notes
		Users
	}
)
