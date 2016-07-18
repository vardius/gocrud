package env

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vardius/goapi"
	"github.com/vardius/golog"
)

var (
	Log    = golog.New()
	Server = goapi.New()
	DB     = connectToDB("root:password@tcp(localhost:3306)/test")
)

func connectToDB(dbURL string) *sql.DB {
	conn, err := sql.Open("mysql", dbURL)
	if err != nil {
		Log.Critical(context.TODO(), "%s", err)
	}

	err = conn.Ping()
	if err != nil {
		Log.Critical(context.TODO(), "%s", err)
	}

	return conn
}
