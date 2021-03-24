package datapool

import "os"

type TargetConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Path     string
	Notify   string
}

func GetTargetConfig() TargetConfig {
	target := TargetConfig{
		Host:     os.Getenv("CRAWLAB_TARGET_HOST"),
		Port:     os.Getenv("CRAWLAB_TARGET_PORT"),
		Username: os.Getenv("CRAWLAB_TARGET_USERNAME"),
		Password: os.Getenv("CRAWLAB_TARGET_PASSWORD"),
		Path:     os.Getenv("CRAWLAB_TARGET_PATH"),
		Notify:   os.Getenv("CRAWLAB_TARGET_NOTIFY"),
	}
	return target
}
