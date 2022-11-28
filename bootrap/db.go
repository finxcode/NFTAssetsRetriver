package bootrap

import (
	"NftAssetsRetriever/global"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func InitDb() (*sql.DB, error) {
	driver := global.App.Config.Database.Driver
	userName := global.App.Config.Database.UserName
	password := global.App.Config.Database.Password
	host := global.App.Config.Database.Host
	port := global.App.Config.Database.Port
	database := global.App.Config.Database.Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", userName, password, host, strconv.Itoa(port), database)

	db, err := sql.Open(driver, dsn)

	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
