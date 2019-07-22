package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var (
	configFilePath = "/Users/sherlock/Desktop/my_project/image_upload/src/image_upload/config/config.yaml"
	ConfigInstance ConfigType
)

// 配置结构
type ConfigType struct {
	IsDebug bool `yaml:"IsDebug"`
	IsLocDev bool `yaml:IsLocDev"`
	BaseDir string `yaml:"BaseDir"`
	// mysql
	MysqlConfig string `yaml:"MysqlConfig"`
	// 日志
	LogFilePath string `yaml:"LogFilePath"`
	// 监听地址
	RestServerAddr string `yaml:"RestServerAddr"`
	// 操作秘钥
	OperationPassword string `yaml:"OperationPassword"`
	// 图片服务配置
	// 图片存放文件夹
	ImageSourceDirectory string `yaml:"ImageSourceDirectory"`
	// 限制图片大小
	ImageSizeLimit int64 `yaml:"ImageSizeLimit"`
	// 图片服务器
	ImageServerHost string `yaml:"ImageServerHost"`
}

func init(){
	configFileEnvPath := os.Getenv("CONFIG_FILE_PATH")
	if configFileEnvPath != ""{
		configFilePath = configFileEnvPath
	}
	fileData, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(fileData, &ConfigInstance)
	if err != nil{
		panic(err)
	}
}
