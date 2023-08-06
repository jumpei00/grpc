package infrastructure

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jumpei00/grpc/grpcback/pkg/config"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func NewMySQL() (*sql.DB, error) {
	dataSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.DataBaseUserName,
		config.DataBasePassword,
		config.DataBaseHost,
		config.DataBasePort,
		config.DataBaseName,
	)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}

	boil.DebugMode = true

	return db, nil
}
