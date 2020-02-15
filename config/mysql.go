//============================================================
// 描述:
// 作者: Yang
// 日期: 2020/2/15 17:37 上午
// 版权: 山东深链智能科技有限公司 @Since 2019
//
//============================================================
package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB



func InitMysqlConnect() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", Conf.Mysql.User, Conf.Mysql.Passwd, Conf.Mysql.Addr, Conf.Mysql.Port, Conf.Mysql.DBName)
	var err error
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(fmt.Errorf("mysql connect err:%v", err, err))
	}

	// 连接池的空闲大小
	DB.SetMaxIdleConns(1000)
	// 最大打开连接数
	DB.SetMaxOpenConns(2000)

	// 测试是否连接成功
	if err := DB.Ping(); err != nil {
		return
	}
}

func CloseDb() {
	defer DB.Close()
}
