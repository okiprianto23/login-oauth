package serverconfig

import (
	"database/sql"
	"github.com/go-redis/redis/v7"
	"github.com/okiprianto23/login-oauth/config"
)

type serverAttribute struct {
	config       config.Configuration
	DBConnection *sql.DB
	redis        *redis.Client
}
