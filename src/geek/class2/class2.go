package class2

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	xerrors "github.com/pkg/errors"
)

var CREATETABLESQL = `
		CREATE TABLE IF NOT EXISTS t_test (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name1 VARCHAR(64) NULL,
			name2 VARCHAR(64) NULL
		);
	`

// 初始化sqlite
func initSqlite() *sql.DB {
	db, _ := sql.Open("sqlite3", "./test.db")
	db.Exec(CREATETABLESQL)
	return db
}

// 清除临时数据
func delSqlite(db *sql.DB) {
	db.Exec("TRUNCATE TABLE t_test")
	db.Close()
	os.Remove("./test.db")
}

// 通过id查找name
func queryName1ById(db *sql.DB, id int) (string, error) {
	var querySQL = "SELECT name1 from t_test where id = ?"
	var retName1 string

	err := db.QueryRow(querySQL, id).Scan(&retName1)
	if err != nil {
		// 自定义错误内容
		return "", xerrors.Wrapf(err, fmt.Sprintf("[queryName1ById] 信息查询失败, id=%v\n", id))
	}
	return retName1, nil
}

// 执行函数
func Run() {
	db := initSqlite()
	defer delSqlite(db)

	name1, err := queryName1ById(db, 1)
	if err != nil {
		// 错误详细信息输出
		fmt.Printf("错误信息为:\n%+v\n", err)

		// 根据业务需要执行对应的错误查看
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("\n\n业务侧对应的处理->没有查询到数据")
		}
		return
	}
	fmt.Printf("this name1=%v\n", name1)

}
