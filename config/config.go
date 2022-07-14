package config

import (
	"BlogProject/cache"
	"BlogProject/logger"
	"BlogProject/model"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	HttpPort    string
	DbHost      string
	DbPort      string
	DbUser      string
	DbPassWord  string
	DbName      string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
)

func Init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		logger.Logger.Info("配置文件读取错误，请检查文件路径:", err)
		panic(err)
	}
	loadServer(file)
	loadMysqlData(file)
	loadRedisData(file)
	LoadQiniu(file)
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	model.InitMysqlDatabase(path)
	cache.InitRedis(RedisAddr, RedisPw, RedisDbName)
}

func loadServer(file *ini.File) {
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func loadMysqlData(file *ini.File) {
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func loadRedisData(file *ini.File) {
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}

func LoadQiniu(file *ini.File)  {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuServer = file.Section("qiniu").Key("QiniuServer").String()
}
