package global

import (
	"NftAssetsRetriever/config"
	"database/sql"
)

type application struct {
	Config config.Configuration
	DB     *sql.DB
}

var App = new(application)
