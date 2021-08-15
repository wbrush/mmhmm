package postgres

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/go-pg/pg/v9"
	"github.com/sirupsen/logrus"
	"github.com/wbrush/go-common/db"
	"github.com/wbrush/mmhmm/models"
)

func (d *PgDAO) CreateNote(note *models.Note) (isDuplicate bool, err error) {
	err = d.BaseDB.RunInTransaction(func(tx *pg.Tx) error {
		//  see if author exists
		_, found, err := d.GetUserById(note.Author.Id)

		//  if not, add them
		if !found || note.Author.Id <= 0 {
			logrus.Debugf("adding author %s %s", note.Author.FirstName, note.Author.LastName)
			_, err = d.CreateUser(note.Author)
			if err != nil {
				return errors.New("Error Creating User: " + err.Error())
			}
		}

		note.UserId = note.Author.Id
		err = tx.Insert(note)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		if db.CheckIfDuplicateError(err) {
			return true, nil
		}
		return false, err
	}

	return false, nil
}

func (d *PgDAO) GetNoteById(id int) (note *models.Note, isFound bool, err error) {
	err = d.BaseDB.RunInTransaction(func(tx *pg.Tx) error {
		//  get note
		note = &models.Note{Id: id}
		err = d.BaseDB.Model(note).WherePK().Select()
		if err != nil {
			return err
		}

		//  get author
		note.Author = &models.User{Id: note.UserId}
		err = d.BaseDB.Model(note.Author).WherePK().Select()
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, false, nil
		}
		return nil, false, err
	}

	return note, true, nil
}

// ListTemplates returns empty array if nothing was found
func (d *PgDAO) ListNotes(filters url.Values) (list []models.Note, err error) {
	list = make([]models.Note, 0)

	//  there is a cleaner way to do this using functions and q.Apply() below. But this way works for a single case
	authorFilter, authorOk := filters["authorId"]
	desiredAuthor := 0
	if authorOk {
		if len(authorFilter) <= 0 {
			authorOk = false
		} else {
			//  set up filter
			desiredAuthor, err = strconv.Atoi(authorFilter[0])
			if err != nil {
				return list, err
			}
		}

		//  delete this key
		delete(filters, "authorId")
	}

	err = d.BaseDB.RunInTransaction(func(tx *pg.Tx) error {
		q := d.BaseDB.Model(&list)
		if authorOk {
			q = q.Where("user_id = ?", desiredAuthor)
		}
		err = q.Select()
		if err != nil && err != pg.ErrNoRows {
			return err
		}

		//  get authors
		idList := make([]int, 0)
		for i := range list {
			idList = append(idList, list[i].UserId)
		}
		authors := make([]models.User, 0)
		err = d.BaseDB.Model(&authors).Where("id in (?)", pg.In(idList)).Select()
		if err != nil {
			return err
		}

		authorsMap := make(map[int]models.User)
		for _, author := range authors {
			authorsMap[author.Id] = author
		}
		for i := range list {
			newAuthor := authorsMap[list[i].UserId]
			list[i].Author = &newAuthor
		}

		return nil
	})
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return list, nil
}

func (d *PgDAO) UpdateNoteById(note *models.Note) (err error) {
	//  all they can do here is edit the text in the note! Until the datamodels or requirements change
	_, err = d.BaseDB.Model(note).WherePK().Set("note = ?", note.Note).Update()
	if err != nil {
		return err
	}

	return nil
}

func (d *PgDAO) DeleteNoteById(id int) (isFound bool, err error) {
	note := models.Note{Id: id}
	res, err := d.BaseDB.Model(&note).WherePK().Delete()
	if err != nil {
		return false, err
	}
	if res.RowsAffected() == 0 {
		return false, nil
	}

	return true, nil
}
