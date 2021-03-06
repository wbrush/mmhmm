package postgres

import (
	"errors"
	"net/url"

	"github.com/go-pg/pg"
	"github.com/wbrush/go-common/db"
	"github.com/wbrush/mmhmm/models"
)

func (d *PgDAO) CreateUser(user *models.User) (isDuplicate bool, err error) {

	err = d.BaseDB.Insert(user)
	if err != nil {
		if db.CheckIfDuplicateError(err) {
			return true, nil
		}
		return false, err
	}

	return false, nil
}

func (d *PgDAO) GetUserById(id int) (user *models.User, isFound bool, err error) {
	user = &models.User{Id: id}
	err = d.BaseDB.Select(user)
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, false, nil
		}
		return nil, false, err
	}

	return user, true, nil
}

// ListTemplates returns empty array if nothing was found
func (d *PgDAO) ListUsers(filters url.Values) (list []models.User, err error) {
	list = make([]models.User, 0)

	err = d.BaseDB.Model(&list).Select()
	if err != nil && err != pg.ErrNoRows {
		return list, err
	}

	return list, nil
}

func (d *PgDAO) UpdateUserById(user *models.User) (err error) {
	result, err := d.BaseDB.Model(user).WherePK().Update()
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		err = errors.New("No record found with given id!")
		return err
	}

	return nil
}

func (d *PgDAO) DeleteUserById(id int) (isFound bool, err error) {
	user := models.User{Id: id}
	result, err := d.BaseDB.Model(&user).WherePK().Delete()
	if err != nil {
		return false, err
	}
	if result.RowsAffected() == 0 {
		return false, nil
	}

	return true, nil
}
