package db

import (
	"blog/model"
)

// AddLeave 添加留言
func AddLeave(leave *model.AddLeave) (id int64, err error) {
	sql := "insert into `leave` (user_name,email,content) values (?,?,?)"
	exec, err := db.Exec(sql, leave.UserName, leave.Email, leave.Content)
	if err != nil {
		return 0, err
	}
	id, err = exec.LastInsertId()
	if err != nil {
		return
	}
	return
}
