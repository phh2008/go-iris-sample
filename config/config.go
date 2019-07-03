package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"reflect"
	"sync"
)

var (
	once sync.Once
	cf   *Conf
)

type Conf struct {
	//根路径
	ContextPath string `yaml:"context-path"`
	//主机端口
	ServerAddress string `yaml:"server-address"`
	//静态资源路径
	StaticPath string `yaml:"static-path"`
}

//初始化
func init() {
	Cfg()
}

//获取配置(单例模式)
func Cfg() *Conf {
	once.Do(func() {
		filePath := "./resources/conf.yml"
		yamlFile, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		cf = new(Conf)
		err = yaml.Unmarshal(yamlFile, cf)
		//打印配置项
		t := reflect.TypeOf(*cf)
		v := reflect.ValueOf(*cf)
		log.Println("-------------------------------------------------------------------")
		for i := 0; i < t.NumField(); i++ {
			log.Printf("|-- %30s  --|--  %20v --|\n", t.Field(i).Name, v.Field(i).Interface())
		}
		log.Println("-------------------------------------------------------------------")
	})
	return cf
}
