package config

import (
	"os"
	"sync"

	"github.com/spf13/viper"
)

type config struct {
	Database Pgsql `yaml:"Pgsql"`
	Mqtt     Mqtt  `yaml:"Mqtt"`
}

type Pgsql struct {
	Host     string `yaml:"Host"`
	DB       string `yaml:"DB"`
	Port     int    `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
}

type Mqtt struct {
	Protocol string `yaml:"Protocol"`
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
}

var once sync.Once

var cfg *config

func Init() *config {
	once.Do(func() {

		//获取项目的执行路径
		path, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		//设置读取的文件路径
		viper.AddConfigPath(path)
		//设置文件的类型
		viper.SetConfigType("yaml")
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
		if err := viper.Unmarshal(&cfg); err != nil {
			panic(err)
		}
	})

	return cfg
}
