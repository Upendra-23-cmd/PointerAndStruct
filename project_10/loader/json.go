package loader

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct{
	AppName	string `json:"app_name"`
	Debug   bool `json:"debug"`

	Server struct{
		Host string `json:"host"`
		Port int	`json:"port"`
	}`json:"server"`
}

func Reader_jsonconfig(filename string)(Config, error){
	var cfg Config

	// Step 1 : Read the file
	data , err := os.ReadFile(filename)
	if err != nil {
		return cfg, fmt.Errorf("failed to read the file : %w", err)
	}

	// parse the json into struct
	if err := json.Unmarshal(data , &cfg); err != nil {
		return cfg, fmt.Errorf("failed to parse the json file : %w", err)
	}
	return cfg, nil
}
