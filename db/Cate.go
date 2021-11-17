package db

import (
	"blog/model"
	"errors"
	"github.com/jmoiron/sqlx"
)

//AddCate 添加分类
func AddCate(cateName string, cateSort int64) (id int64, err error) {
	sql1 := `select id,cate_name,cate_sort from category where cate_name=?`
	var cateInfo []*model.Cate
	err = db.Select(&cateInfo, sql1, cateName)
	if err != nil {
		return
	}
	if cateInfo != nil {
		err = errors.New("分类已经存在")
		return
	}

	sql := `insert into category (cate_name,cate_sort) values(?,?)`
	res, err := db.Exec(sql, cateName, cateSort)
	if err != nil {
		return
	}
	id, err = res.LastInsertId()
	return
}

//AllCate 获取全部分类
func AllCate() (list map[int64]interface{}, err error) {
	sql1 := `select id,cate_name,cate_sort from category`
	var cateInfo []*model.Cate
	err = db.Select(&cateInfo, sql1)
	if err != nil {
		return
	}
	type cateDetail struct {
		Name string
		Sort int64
	}
	if cateInfo == nil {
		return
	}
	list = make(map[int64]interface{})
	for _, v := range cateInfo {
		var tmp cateDetail
		tmp.Sort = v.CateSort
		tmp.Name = v.CateName
		list[v.Id] = tmp
	}
	return
}

//GetCateByIds 通过id查询分类
func GetCateByIds(ids []int64) (list map[int64]string, err error) {
	sql := `select id,cate_name,cate_sort from category where id in(?)`
	query, args, err := sqlx.In(sql, ids)
	if err != nil {
		return
	}
	var cateInfo []*model.Cate
	err = db.Select(&cateInfo, query, args...)
	if err != nil {
		return
	}
	list = make(map[int64]string)
	for _, v := range cateInfo {
		list[v.Id] = v.CateName
	}
	return
}
