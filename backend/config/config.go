package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App    AppConfig    `mapstructure:"app"`
	Logger LoggerConfig `mapstructure:"logger"`
	DB     DBConfig     `mapstructure:"db"`
}

type AppConfig struct {
	Name     string `mapstructure:"name"`
	Version  string `mapstructure:"version"`
	HttpPort int    `mapstructure:"httpPort"`
	GrpcPort int    `mapstructure:"grpcPort"`
}

type LoggerConfig struct {
	IsEnable bool `mapstructure:"isEnable"`
}

type DBConfig struct {
	Driver             string `mapstructure:"driver"`
	Username           string `mapstructure:"username"`
	Password           string `mapstructure:"password"`
	Name               string `mapstructure:"name"`
	Host               string `mapstructure:"host"`
	Port               int    `mapstructure:"port"`
	MaxIdleConnections int    `mapstructure:"maxIdleConnections"`
	MaxOpenConnections int    `mapstructure:"maxOpenConnections"`
	DebugMode          bool   `mapstructure:"debugMode"`
	Timeout            string `mapstructure:"timeout"`
	WriteTimeout       string `mapstructure:"writeTimeout"`
	ReadTimeout        string `mapstructure:"readTimeout"`
	SSLMode            string `mapstructure:"sslmode"`
}

func Load() Config {
	viper.SetConfigFile("config")
	viper.SetConfigFile("./resources/config.json")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	conf := Config{}
	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}

	return conf
}

func (db *DBConfig) GetDSN() (dsn string) {
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%s&writeTimeout=%s&readTimeout=%s&charset=utf8mb4&parseTime=True&loc=Local",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
		db.Timeout,
		db.WriteTimeout,
		db.ReadTimeout,
	)

	if db.Driver == "postgres" {
		dsn = fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s",
			db.Driver,
			db.Username,
			db.Password,
			db.Host,
			db.Port,
			db.Name,
			db.SSLMode,
		)
	}

	fmt.Println(dsn)

	return
}
