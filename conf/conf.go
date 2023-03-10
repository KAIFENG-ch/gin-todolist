package conf

import (
	"gin-todolist/cache"
	"gin-todolist/model"
	"gin-todolist/pkg/util"
	"gopkg.in/ini.v1"
	"log"
	"strings"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func Init() {
	file, err := ini.Load("conf/config.ini")
	if err != nil {
		util.LogrusObj.Info("配置文件读取错误，请检查文件路径:", err)
		log.Fatal(err)
	}
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.LogrusObj.Info(err) //日志内容
		log.Fatal(err)
	}
	LoadServer(file)
	LoadMysqlData(file)
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(path)
	cache.Redis()
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
