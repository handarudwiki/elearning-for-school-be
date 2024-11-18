package config

type Config struct {
	Server   Server
	Database Database
	JWT      JWT
}

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type JWT struct {
	AccessSecretKey  string
	RefreshSecretKey string
}
