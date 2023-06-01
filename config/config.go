package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var Conf Config

func InitConfig() {
	dataBytes, err := os.ReadFile("config.yaml")
	if err != nil {
		panic("读取文件失败：" + err.Error())
	}
	Conf := Config{}
	err = yaml.Unmarshal(dataBytes, &Conf)
	if err != nil {
		panic("解析 yaml 文件失败：" + err.Error())
	}
	fmt.Printf("config → %+v\n", Conf)

}

type Config struct {
	Mysql Mysql `json:"mysql"`
	// Redis Redis `json:"redis"`
	LogConfig LogConfig `json:"logconfig"`
}
type Mysql struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

type LogConfig struct {
	Dir        string `yaml:"dir"`
	Name       string `yaml:"name"`
	Level      string `yaml:"level"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
}
