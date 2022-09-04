package config

import (
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

var (
	TestDB = DB{
		Host:   "192.168.50.200",
		Port:   3306,
		User:   "root",
		Pass:   "123456",
		DBName: "im",
	}
)

var (
	c = config{}
)

func GetWeb() Web {
	return c.Web
}

func GetDB() DB {
	return c.DB
}

type config struct {
	Web Web `yaml:"web"`
	DB  DB  `yaml:"db"`
}

type Web struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type DB struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	DBName string `yaml:"db_name"`
}

func Init(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		logrus.Fatalf("open %s err: %s", filename, err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		logrus.Fatalf("decode %s err: %s", filename, err)
	}
}
