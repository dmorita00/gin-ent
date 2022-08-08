package configutil

import (
	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/name/pjName/server/ent"
	"golang.org/x/xerrors"
)

type MySQLConfig struct {
	DSN string `envconfig:"DSN" default:""`

	// 発行されるSQLをロギングするかどうか
	Debug bool `envconfig:"MYSQL_DEBUG" default:"false"`
}

var (
	db    *ent.Client
	rawDB *sql.Driver
)

func DBInit() error {
	sc := GetServerConfig()
	entOptions := []ent.Option{}
	if sc.MySQL.Debug {
		entOptions = append(entOptions, ent.Debug())
	}

	rawDB, err := sql.Open("mysql", sc.MySQL.DSN)
	if err != nil {
		return xerrors.Errorf("message: %w", err)
	}

	db = ent.NewClient(append(entOptions, ent.Driver(rawDB))...)

	return nil
}

func DBClose() error {
	err := db.Close()
	if err != nil {
		return xerrors.Errorf("message: %w", err)
	}
	return nil
}

func GetDB() *ent.Client {
	return db
}

// GetRawDB 生SQLを実行するためのdriverを取得する
func GetRawDB() *sql.Driver {
	return rawDB
}
