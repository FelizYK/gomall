package conf

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/kr/pretty"
	"gopkg.in/yaml.v3"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Env string `yaml:"env" validate:"required"`

	Redis Redis `yaml:"redis"`
	Etcd  Etcd  `yaml:"etcd"`
}

type Redis struct {
	Address string `yaml:"address" validate:"required"`
	DB      int    `yaml:"db" validate:"min=0"`
}

type Etcd struct {
	Endpoints string `yaml:"endpoints" validate:"required"`
}

// GetConf gets configuration instance
func GetConf() *Config {
	once.Do(initConf)
	return conf
}

func initConf() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// load conf.yaml
	confFileRelPath := filepath.Join("conf", filepath.Join(getEnv(), "conf.yaml"))
	content, err := os.ReadFile(confFileRelPath)
	if err != nil {
		panic(err)
	}

	conf = new(Config)
	conf.Env = getEnv()
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		panic(err)
	}

	validate := validator.New()
	if err := validate.Struct(conf); err != nil {
		panic(err)
	}
	pretty.Printf("%+v\n", conf)
}

func getEnv() string {
	e := os.Getenv("GO_ENV")
	if len(e) == 0 {
		return "dev"
	}
	return e
}
