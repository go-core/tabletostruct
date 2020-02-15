//============================================================
// 描述:
// 作者: Yang
// 日期: 2020/2/15 16:39 上午
// 版权: 山东深链智能科技有限公司 @Since 2019
//
//============================================================
package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	Conf     *config
	confPath string //配置文件路径
)

func init() {
	flag.StringVar(&confPath, "c", "./app.yaml", "配置文件")
}

type config struct {
	Mysql *MysqlConf `yaml:"mysql"`
}

type MysqlConf struct {
	Addr    string `yaml:"addr"`
	Port    string `yaml:"port"`
	User    string `yaml:"user"`
	Passwd  string `yaml:"passwd"`
	DBName  string `yaml:"dbname"`
	Exclude string `yaml:"exclude"`
}

func InitConfigration() {
	hs, err := ioutil.ReadFile(confPath)
	if err != nil {
		panic("配置文件读取失败")
	}
	err = yaml.Unmarshal(hs, &Conf)
	if err != nil {
		fmt.Printf("配置文件数据序列化失败%v", err.Error())
	}

}
