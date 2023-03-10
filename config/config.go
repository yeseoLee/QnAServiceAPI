package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Service service `yaml:"service"`
		DB      db      `yaml:"db"`
	}
	service struct {
		Port string `yaml:"port"`
		Root string `yaml:"root"`
	}
	db struct {
		DBType string `yaml:"dbType"`
		Sqlite sqlite `yaml:"sqlite"`
		Mysql  mysql  `yaml:"mysql"`
	}
	mysql struct {
		Host         string `yaml:"host"`
		Port         int    `yaml:"port"`
		DatabaseName string `yaml:"databaseName"`
		User         string `yaml:"user"`
		Password     string `yaml:"password"`
	}
	sqlite struct {
		FilePath string `yaml:"filePath"`
	}
)

var instance Config

func GetInstance() Config {
	return instance
}

func init() {
	profile := initProfile()
	setRuntimeConfig(profile)
}

// 환경변수로 Config용 PROFILE 가져오기
func initProfile() string {
	var profile string
	profile = os.Getenv("GO_PROFILE")
	if len(profile) <= 0 {
		profile = "config_local"
	}
	fmt.Println("GOLANG_PROFILE: " + profile)
	return profile
}

func setRuntimeConfig(profile string) {
	viper.AddConfigPath("./config/")
	viper.SetConfigName(profile)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&instance)
	if err != nil {
		log.Fatal(err)
	}

	/*
		// 설정파일이 변경되면 이벤트를 핸들링하여 다시 언마샬링
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("Config file changed:", e.Name)
			var err error
			err = viper.ReadInConfig()
			if err != nil {
				fmt.Println(err)
				return
			}
			err = viper.Unmarshal(&Conf)
			if err != nil {
				fmt.Println(err)
				return
			}
		})
		viper.WatchConfig()
	*/
}

// 사용예시: Conf.DB.DBType
