package serverconfig

import (
	"database/sql"
	"github.com/go-redis/redis/v7"
	"github.com/okiprianto23/common/http/client"
	"github.com/okiprianto23/common/util/validator/basic_validator"
	"github.com/okiprianto23/common/util/validator/tag_validator"
	"github.com/okiprianto23/login-oauth/config"
)

type serverAttribute struct {
	config       config.Configuration
	DBConnection *sql.DB
	redis        *redis.Client

	//internalTokenValidator token.InternalJWTValidator
}

type validator struct {
	basicValidator basic_validator.BasicValidator
	tagValidator   tag_validator.TagValidator
	//listDataValidator lisdatava
	//HttpController *httppcon
	apiConnector client.APIConnector
}

type services struct {
}

type listDAO struct {
}
