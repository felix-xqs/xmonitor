package conf

import "github.com/felix-xqs/golog"

type Conf struct {
	Port     int    `yaml:"port"`
	RootPath string `yaml:"rootPath"`
	Environ  string `yaml:"environ"`
	Svrs     map[string]string
	Log      *golog.Config
}
