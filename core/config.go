package core

import (
	"fast_gin/config"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const configFile = "settings.yaml"

// ConfInit 读取yaml配置
func ConfInit() config.Server {

	c := &config.Server{}
	yamlConf, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}

	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		logrus.Fatalf("config Init Unmarshal: %v", err)
	}
	logrus.Println("config yamlFile load Init success.")
	return *c
}
