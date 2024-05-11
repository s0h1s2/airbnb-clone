package config

import "fmt"

type jwtConfig struct {
	JwtSecretKey string `env:"JWT_SECRET_KEY"`
}
type databaseConfig struct {
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
	Host     string `env:"DB_HOST"`
	Port     uint16 `env:"DB_PORT"`
}
type config struct {
	Jwt           jwtConfig
	Db            databaseConfig
	ForntendProxy string `env:"PROXY_FRONTEND"`
}

var Config config

func (db databaseConfig) String() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s", db.User, db.Password, db.Host, db.Port, db.Name)
}
