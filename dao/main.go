package dao

import (
	"net/url"

	"github.com/wbrush/go-common/db"
	"github.com/wbrush/mmhmm/models"
)

type (
	Notes interface {
		// CreateTemplate(shardID int64, template *datamodels.Template) (isDuplicate bool, err error)
		// GetTemplateById(shardID int64, id int64) (template *datamodels.Template, isFound bool, err error)
		// ListTemplates(shardID int64, filters url.Values) (templates []datamodels.Template, total int, hasMore bool, err error)
		// UpdateTemplate(shardID int64, template *datamodels.Template) (err error)
		// DeleteTemplateById(shardID int64, id int64) (isFound bool, err error)
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
