package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config 是一个全局配置，会有Database,server以及其他的的配置
type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	Server   ServerConfig   `mapstructure:"server"`
}
type DatabaseConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	DBName    string `mapstructure:"dbname"`
	Charset   string `mapstructure:"charset"`
	ParseTime bool   `mapstructure:"parseTime"`
	Loc       string `mapstructure:"loc"` //使用mapstructure的目的是与yaml中的键对应
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

var GlobalConfig Config

func InitConfig() error {
	//设置viper读取的文件名，类型以及地址
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	//读取config.yaml的内容，并将其按照mapstructure的tag赋值给GlobalConfig
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %s", err)
	}
	err = viper.Unmarshal(&GlobalConfig)
	if err != nil {
		return fmt.Errorf("解析配置文件失败: %s", err)
	}

	return nil
}
