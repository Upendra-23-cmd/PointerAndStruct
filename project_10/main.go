package main

import (
	"fmt"
	"project_10/loader"
)

func loader_config() {

	// Load config.json file
	cfg, err := loader.Reader_jsonconfig("config.json")
	if err != nil {
		panic(err)
	}

	fmt.Println("App Name   :", cfg.AppName)
	fmt.Println("Debug Mode :", cfg.Debug)
	fmt.Println("Server Host:", cfg.Server.Host)
	fmt.Println("Server Port:", cfg.Server.Port)
}

func yaml_loader_file() {

	cfg, err := loader.Reader_yamlconfig("config.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Println("App Name   :", cfg.AppName)
	fmt.Println("Debug Mode :", cfg.Debug)
	fmt.Println("Server Host:", cfg.Server.Host)
	fmt.Println("Server Port:", cfg.Server.Port)
}

func main() {
	yaml_loader_file()
	fmt.Println("==========================================")
	loader_config()
}
