package db

import (
	Mydb "fileStoreServer/db/mysql"
	"fmt"
)

//文件上传完成
func MysqlUpdateFile(filehash string, filename string, filesize int64, fileArr string) bool {
	prepare, err := Mydb.DBConn().Prepare(
		"insert  into tbl_file('file_sha1','file_name','file_size','file_addr','file_addr')values (?,?,?,?,1)")

	defer prepare.Close()
	if err != nil {
		fmt.Println("error is %s", err.Error())
		return false
	}
	exec, err := prepare.Exec(filehash, filename, filesize, fileArr)

	if err != nil {
		fmt.Println("sql 解析错误%s", err.Error())
		return false
	}
	if rf, err := exec.RowsAffected(); err == nil {
		if rf <= 0 {
			fmt.Println("warning nothind is been upload ,err=%s", err.Error())
		}
		return true
	}
	return false
}
