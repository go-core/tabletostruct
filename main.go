//============================================================
// 描述:
// 作者: Yang
// 日期: 2020/2/15 16:39 上午
// 版权: 山东深链智能科技有限公司 @Since 2019
//
//============================================================
package main

import (
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"strings"
	"tabletostruct/config"
	"unicode"
)

//数据字段类型
type Column struct {
	Columname     string `json:"columname"`
	Datatype      string `json:"datatype"`
	Columncomment string `json:"columncomment"`
	Columnkey     string `json:"columnkey"`
	Extra         string `json:"extra"`
}

func main() {
	//加载配置文件
	config.InitConfigration()
	config.InitMysqlConnect()
	//命令行输入参数
	var t string
	flag.StringVar(&t, "t", "", "填写上要生成的数据表名称")
	flag.Parse()

	//获取要排除的字段
	str := getExcludeStr()
	//如果字段为空
	var cloumn_sql string
	if len(str) == 0 {
		cloumn_sql += fmt.Sprintf("select column_name columnName, data_type dataType, column_comment columnComment, column_key columnKey, extra from information_schema.columns where table_name = '%s'", t)
	} else {
		cloumn_sql += fmt.Sprintf("select column_name columnName, data_type dataType, column_comment columnComment, column_key columnKey, extra from information_schema.columns where table_name = '%s' and column_name not in(%s)", t, str)
	}
	//获取所有字段
	cloumns, err := config.DB.Query(cloumn_sql)
	if err != nil {
		panic(fmt.Errorf("查询失败:%v", err.Error()))
	}

	var struct_str string
	struct_str = fmt.Sprintf("type %s struct {\n", t)

	column := &Column{}
	for cloumns.Next() {
		err := cloumns.Scan(&column.Columname, &column.Datatype, &column.Columncomment, &column.Columnkey, &column.Extra)
		if err != nil {
			fmt.Printf("Scan failed,err:%v", err)
			return
		}

		struct_str += "    " + column.Columname
		//根据字段数据类型判断
		if column.Datatype == "int" {
			struct_str += " libs.Int "
		} else if column.Datatype == "decimal" {
			struct_str += " libs.Float "
		} else if column.Datatype == "tinyint" {
			struct_str += " libs.Int " //需要手工查看并根据实际情况作出修改
		} else if column.Datatype == "timestamp" {
			struct_str += " libs.Int "
		} else if column.Datatype == "datetime" {
			struct_str += "libs.Time"
		} else if column.Datatype == "json" {
			struct_str += "libs.Ext"
		} else {
			struct_str += " libs.String "
		}
		struct_str += fmt.Sprintf("`db:\"%s\" json:\"%s\"` //%s \n", column.Columname, firstCharToLower(column.Columname), column.Columncomment)
	}

	struct_str += "    " + "storages.BaseItem\n"
	struct_str += "}"

	// 复制内容到剪切板
	clipboard.WriteAll(struct_str)
	//关闭数据库连接
	config.CloseDb()

}

func getExcludeStr() string {
	exclude := strings.Split(config.Conf.Mysql.Exclude, ",")
	var excludeStr string
	if len(exclude) > 1 {
		for _, v := range exclude {
			excludeStr += strings.Join([]string{"\"", v, "\"", ","}, "")
		}
	}

	if len(excludeStr) > 0 {
		excludeStr = strings.TrimRight(excludeStr, ",")
	}

	return excludeStr
}

//首字母小写
func firstCharToLower(str string) string {
	var n []rune
	for i, ch := range str {
		if i == 0 {
			n = append(n, unicode.ToLower(ch))
		} else {
			n = append(n, ch)
		}
	}
	return string(n)
}
