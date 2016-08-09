package env

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vardius/golog"
	"github.com/vardius/goserver"
	"golang.org/x/net/context"
)

var (
	Log    golog.Logger
	Server goserver.Server
	DB     *sql.DB
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

func init() {
	Log = golog.New()
	Server = goserver.New()
	DB = connectToDB("root:password@tcp(127.0.0.1:3306)/test")
}
