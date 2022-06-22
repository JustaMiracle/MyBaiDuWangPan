package db

import (
	myDb "fileStoreServer/db/mysql"
	"fmt"
)

//通过账号密码完成user表的注册操作
func UserSignUp(username string, password string) bool {
	//前端传来用户信息，实现登入逻辑
	fmt.Println("你打野的0")
	prepare, err := myDb.DBConn().Prepare(
		"insert  into tbl_user(`user_name`,`user_pwd`)values (?,?)")
	fmt.Println("你打野的1")
	if err != nil {
		fmt.Println("111111111111failed to insert%s", err.Error())
		return false
	}
	fmt.Println("你打野的2")
	defer prepare.Close()
	exec, err := prepare.Exec(username, password)
	fmt.Println("你打野的3")
	if err != nil {
		fmt.Println("failed to insert%s", err.Error())
		return false
	}
	fmt.Println("你打野的4")
	if affected, err := exec.RowsAffected(); err == nil && affected > 0 {
		return true
	}
	fmt.Println("你打野的")
	return false
}
