package config

import (
	"fmt"
	"os" 
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-default:"local" env-reqired: "true"`
	StoragePath string `yaml:"storage_path" `
}
type HTTPServer struct {
	Address string `yaml:"storage_path" `
}

func Mustload() *Config{
	configPath := os.Getenv( "CONFIG_PATH")
	if configPath == "" {
		fmt.Println("configPath is not set")
	}
	if _,err :=os.Stat(configPath);os.IsNotExist(err){
		fmt.Println("config file does not exist")
		return nil
	}
	var cfg Config
	if err:=cleanenv.ReadConfig(configPath,&cfg); err !=nil {
		fmt.Println("config CAN NOT BE RED", err)
		return nil
	}
	return &cfg

}

