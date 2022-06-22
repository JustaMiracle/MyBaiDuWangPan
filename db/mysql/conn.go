package mysql

import (
	"database/sql"
	"fmt"
	"os"
)

var db *sql.DB

//const (
//	userName = "root"
//	password = "123123"
//	ip       = "127.0.0.1"
//	port     = "3306"
//	dbName   = "test1"
//)

func init() {
	db, _ = sql.Open("mysql", "root:123123@(127.0.0.1:3306)/test1") // 设置连接数据库的参数
	fmt.Println("你打野121")
	//	defer db.Close() //关闭数据库
	err := db.Ping() //连接数据库
	if err != nil {
		fmt.Println("数据库连接失败")
		os.Exit(1)
	}

}

//返回数据连接对象
func DBConn() *sql.DB {
	fmt.Println("你打野11111111")
	fmt.Println(db)
	return db
}
