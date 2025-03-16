package configs

import (
	"time"
)

var MAX_DB_OPEN_RETRIES = 5

const (
	TOKEN_LIFETIME          = time.Hour * 24
	JWT_SECRET_FILE_ENV_VAR = "JWT_SECRET_FILE"
)