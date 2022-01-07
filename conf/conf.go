package conf

import (
	"TodoList/model"
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	AppMode    string
	AppHost    string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func Init() {
	// 读取文件配置
	file, err := ini.Load("./conf/config.ini")
	// 抛出panic
	if err != nil {
		fmt.Println("文件未找到")
		panic(err)
	}
	// 加载配置
	LoadServer(file)
	LoadMysql(file)
	// 连接mysql
	path := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=True"}, "")
	model.DataBase(path)
}

// 加载系统配置
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").String()
	AppHost = file.Section("server").Key("AppHost").String()
}

// 加载数据库配置
func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("DataBase").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
