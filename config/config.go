package config

import (
	"flag"
	"github.com/spf13/viper"
)

var  Ops options

func SetUp(){
	var configFile string
	flag.StringVar(&configFile, "c", "", "specified config file")
	flag.Parse()
	if len(configFile) == 0{
		panic("config file not found.")
	}
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil{
		panic(err)
	}
	//viper.WatchConfig()

	err = viper.Unmarshal(&Ops)
	if err != nil{
		panic(err)
	}

}

type options struct {
	Server Server `yaml:"server"`
	Mysql  Mysql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
}
type Server struct {
	Port int `yaml:"port"`
}
type Mysql struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	User            string `yaml:"user"`
	Passwd          string `yaml:"passwd"`
	Db              string `yaml:"db"`
	MaxConnects     int    `yaml:"max_connects"`
	MaxIdleConnects int    `yaml:"max_idle_connects"`
	ConnMaxIdleTime int    `yaml:"conn_max_idle_time"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}