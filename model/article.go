package model

type Article struct {
	Id           int    `db:"id"`
	Email        string `db:"email"`
	UserName     string `db:"user_name"`
	CateId       int    `db:"cate_id"`
	Title        string `db:"title"`
	Summary      string `db:"summary"`
	ViewCount    int    `db:"view_count"`
	CommentCount int    `db:"comment_count"`
	CreateTime   string `db:"create_time"`
	CateName     string
}

type ArticleDetail struct {
	Article
	Content string `db:"content"`
}
