package postgres

import (
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

// func (d *PgDAO) GetTemplateById(shardId int64, id int64) (template *datamodels.Template, isFound bool, err error) {
// 	if !d.ValidateCluster(shardId) {
// 		err = errors.New("cluster is not ready yet")
// 		return
// 	}

// 	template = &datamodels.Template{Id: id}
// 	err = d.Cluster.Shard(shardId).Select(template)
// 	if err != nil {
// 		if err == pg.ErrNoRows {
// 			return nil, false, nil
// 		}
// 		return nil, false, err
// 	}

// 	selfPath := d.buildSelfPath(template.Id)

// 	//TODO really need this?
// 	if !strings.EqualFold(template.TemplateSelf, selfPath) {
// 		//means self path changed
// 		template.TemplateSelf = selfPath
// 		err = d.Cluster.Shard(shardId).Update(template)
// 		if err != nil {
// 			return template, true, err
// 		}
// 	}

// 	return template, true, nil
// }

// // ListTemplates returns empty array if nothing was found
// func (d *PgDAO) ListTemplates(shardId int64, filters url.Values) (list []datamodels.Template, total int, hasMore bool, err error) {
// 	if !d.ValidateCluster(shardId) {
// 		err = errors.New("cluster is not ready yet")
// 		return
// 	}

// 	list = make([]datamodels.Template, 0)
// 	hasMore = false

// 	pf, err := db.PrepareFiltersByModel(filters, datamodels.Template{})
// 	if err != nil {
// 		return list, 0, hasMore, err
// 	}
// 	//also check unknown fields errors
// 	if len(pf.Errors) > 0 {
// 		return list, 0, hasMore, fmt.Errorf("%v", pf.Errors)
// 	}

// 	f := new(datamodels.TemplateFilter)
// 	err = urlstruct.Unmarshal(context.Background(), pf.Prepared, f)
// 	if err != nil {
// 		return list, 0, hasMore, err
// 	}

// 	q := d.Cluster.Shard(shardId).Model(&list).
// 		WhereStruct(f).
// 		Limit(f.Pager.Limit).
// 		Offset(f.Pager.GetOffset())

// 	q, err = db.ApplyDefaultFilters(q, pf.Prepared)
// 	if err != nil {
// 		return list, 0, hasMore, err
// 	}

// 	total, err = q.SelectAndCount()
// 	if err != nil && err != pg.ErrNoRows {
// 		return list, total, hasMore, err
// 	}

// 	if len(list) > 0 {
// 		hasMore, err = q.Where("?TableAlias.id > ?", list[len(list)-1].Id).Exists()
// 		if err != nil && err != pg.ErrNoRows {
// 			return list, total, hasMore, err
// 		}
// 	}

// 	return list, total, hasMore, nil
// }

// func (d *PgDAO) UpdateTemplate(shardId int64, template *datamodels.Template) (err error) {
// 	if !d.ValidateCluster(shardId) {
// 		err = errors.New("cluster is not ready yet")
// 		return
// 	}

// 	err = d.Cluster.Shard(shardId).Update(template)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (d *PgDAO) DeleteTemplateById(shardId int64, id int64) (isFound bool, err error) {
// 	if !d.ValidateCluster(shardId) {
// 		err = errors.New("cluster is not ready yet")
// 		return
// 	}

// 	template := datamodels.Template{Id: id}
// 	res, err := d.Cluster.Shard(shardId).Model(&template).WherePK().Delete()
// 	if err != nil {
// 		return false, err
// 	}
// 	if res.RowsAffected() == 0 {
// 		return false, nil
// 	}

// 	return true, nil
// }
