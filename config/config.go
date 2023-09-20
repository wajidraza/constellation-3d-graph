package config

import "os"

type Config struct {
	AppName string
	Port    string
	DBUrl   string
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return &Config{
		AppName: "Constellation 3D Knowledge Graph Visualizer",
		Port:    port,
		DBUrl:   os.Getenv("DATABASE_URL"),
	}
}
