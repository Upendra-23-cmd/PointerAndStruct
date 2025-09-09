package loader

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v3"
)

type Config_Yaml struct{
	AppName	string `yaml:"app_name"`
	Debug   string `yaml:"debug"`

	Server struct{
		Host string `yaml:"host"`
		Port int	`yaml:"port"`
	}`yaml:"server"`
}

func Reader_yamlconfig(filename string)(Config_Yaml, error){

	var cfg Config_Yaml
	// read the yaml file
	data, err := os.ReadFile(filename)
	if err != nil {
		return cfg, fmt.Errorf("error to load the file : %w",err) 
	}

	// parse the value to struct
	if err := yaml.Unmarshal(data, &cfg);err != nil{
		return cfg, fmt.Errorf("unable to parse the yaml file : %w",err)
	}
	return cfg, nil
}

