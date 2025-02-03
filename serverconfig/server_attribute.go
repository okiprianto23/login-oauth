package serverconfig

import (
	"github.com/okiprianto23/common/connection/database"
	"github.com/okiprianto23/common/connection/redis"
	"github.com/okiprianto23/login-oauth/config"
)

func NewServerAttribute(config config.Configuration) serverAttribute {
	return serverAttribute{config: config}
}

func (s *serverAttribute) Init() (err error) {

	//Connection Database
	s.DBConnection = database.GetDBConnection(
		database.DBAddressParam().
			Host(s.config.PostgresSql.Host).
			Port(s.config.PostgresSql.Port).
			DBName(s.config.PostgresSql.DBName).
			Password(s.config.PostgresSql.Password).
			SSLMode(s.config.PostgresSql.SslMode).
			Username(s.config.PostgresSql.Username).
			MaxIdleConnection(s.config.PostgresSql.MaxIdleConnection).
			MaxOpenConnection(s.config.PostgresSql.MaxOpenConnection).
			DefaultSchema(s.config.PostgresSql.DefaultSchema).
			SetDriver("pgx"),
	)

	//Connection Redis
	s.redis = redis.ConnectRedis(
		redis.NewConnectionRedis(
			s.config.Redis.Host,
			s.config.Redis.Port,
		).
			SetDatabase(s.config.Redis.DB).
			SetMaxRetires(s.config.Redis.MaxRetries).
			SetPassword(s.config.Redis.Password).
			SetUsername(s.config.Redis.Username),
	)

	return
}
