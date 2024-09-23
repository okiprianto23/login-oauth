package config

import (
	"encoding/json"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/okiprianto23/zerolog"
	logzero "github.com/okiprianto23/zerolog/log"
	"io"
	"log"
	"os"
	"time"
)

var (
	AppConfig Configuration
)

type Configuration struct {
	Server      server      `envconfig:"server"`
	PostgresSql PostgresSQL `envconfig:"postgresSql"`
	Redis       redis       `envconfig:"redis"`
	Log         logs        `envconfig:"log"`
	Token       token       `envconfig:"token"`
	Credentials credentials `envconfig:"credentials"`
}

type server struct {
	Application string `envconfig:"application"`
	Version     string `envconfig:"version" default:"1.0.0"`
	Host        string `envconfig:"host"`
	Port        int    `envconfig:"port" default:"8080"`
}

type PostgresSQL struct {
	Host              string `envconfig:"host" default:"localhost"`
	Port              int    `envconfig:"port" default:"5432"`
	Username          string `envconfig:"username"`
	Password          string `envconfig:"password"`
	DBName            string `envconfig:"dbname"`
	DefaultSchema     string `envconfig:"default-schema"`
	SslMode           string `envconfig:"ssl-mode" default:"disable"`
	MaxOpenConnection int    `envconfig:"max_open_connection" default:"50"`
	MaxIdleConnection int    `envconfig:"max_idle_connection" default:"10"`
}

type redis struct {
	Host       string `envconfig:"host" default:"localhost"`
	Port       int    `envconfig:"port"`
	DB         int    `envconfig:"DB"`
	Password   string `envconfig:"password"`
	Username   string `envconfig:"username"`
	MaxRetries int    `envconfig:"max_retries"`
}

type logs struct {
	Level int8 `envconfig:"level" default:"0"`
}

type credentials struct {
	ClientID string `envconfig:"client_id"`
	UserID   string `envconfig:"user_id"`
}

type token struct {
	UserKey            string        `envconfig:"user_key"`
	FixedInternalToken string        `envconfig:"fixed"`
	Duration           time.Duration `envconfig:"duration" default:"24h"`
}

func init() {

	if err := envconfig.Process(
		"oauth",
		&AppConfig,
	); err != nil {
		log.Fatal(err)
	}

	var logwriter []io.Writer
	logwriter = append(logwriter, os.Stderr)

	//TODO for logging publis
	//if AppConfig.Log.pub

	zerolog.InitLoggerModel(AppConfig.Server.Application, AppConfig.Server.Version)
	zerolog.SetGlobalLevel(zerolog.Level(AppConfig.Log.Level))
	zerolog.ErrorHandler = func(err error) {
		log.Println("Something Trouble Found on writing log ", err)
	}

	logger := zerolog.New(zerolog.MultiLevelWriter(logwriter...)).With().Timestamp().Logger()
	logzero.InitiateLogger(logger)

	printConfig(AppConfig)
}

func printConfig(c Configuration) {
	data, _ := json.MarshalIndent(c, "", "\t")
	fmt.Println(string(data))
}
