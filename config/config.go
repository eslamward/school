package config

type AppConfig struct {
	Server
	App
}

type App struct {
	Name string
}
type Server struct {
	Host string
	Port string
}
