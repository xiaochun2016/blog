package model

type Cate struct {
	Id       int64  `db:"id"`
	CateName string `db:"cate_name"`
	CateSort int64  `db:"cate_sort"`
}
