package db

import "blog/model"

// AddArticle 添加文章
func AddArticle(username string, title string, email string, content string, cateId int64, summary string) (id int64, err error) {
	sql := `insert into article (user_name,email,cate_id,title,summary,content) values (?,?,?,?,?,?)`
	exec, err := db.Exec(sql, username, email, cateId, title, summary, content)
	if err != nil {
		return 0, err
	}
	id, err = exec.LastInsertId()
	if err != nil {
		return
	}
	return
}

// ArtList 文件列表
func ArtList(params *model.PostParams) (list []*model.Article, err error) {
	sql1 := `select id,user_name,email,cate_id,title,summary,create_time from article where status=1 limit ?,?`
	err = db.Select(&list, sql1, (params.Page-1)*params.PageSize, params.PageSize)
	if err != nil {
		return
	}
	return
}

//ArtListById 通过分类id获取文章列表
func ArtListById(id, page, pageSize int) (list []*model.Article, err error) {
	sql1 := `select id,user_name,email,cate_id,title,summary,create_time from article where status=1 and cate_id=? limit ?,?`
	err = db.Select(&list, sql1, id, (page-1)*pageSize, pageSize)
	if err != nil {
		return
	}
	return
}
