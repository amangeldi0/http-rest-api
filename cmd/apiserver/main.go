package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/amangeldi0/http-rest-api/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

func inti() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "Path to config file")
}

func main() {
	inti()

	flag.Parse()

	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
