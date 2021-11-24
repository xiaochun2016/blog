package model

type LeaveInfo struct {
	Id         int64  `db:"id"`
	UserName   string `db:"user_name" form:"user_name" binding:"required"`
	Email      string `db:"email" form:"email" binding:"required,email"`
	Content    string `db:"content" form:"content" binding:"required"`
	CreateTime string `db:"create_time"`
}

type AddLeave struct {
	UserName string `db:"user_name" form:"user_name" binding:"required"`
	Email    string `db:"email" form:"email" binding:"required,email"`
	Content  string `db:"content" form:"content" binding:"required"`
}
