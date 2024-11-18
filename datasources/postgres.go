package datasources

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type PostgresSource interface {
	GetConnection() *sqlx.DB
}

type postgresSource struct {
	dbHost     string
	dbName     string
	dbUser     string
	dbPassword string
	dbSslMode  string
}

func NewPostgresSource(dbHost, dbName, dbUser, dbPassword, dbSslMode string) PostgresSource {
	return &postgresSource{dbHost, dbName, dbUser, dbPassword, dbSslMode}
}

func (p *postgresSource) GetConnection() *sqlx.DB {
	dbDriver := "postgres"

	connectionString := "host=" + p.dbHost +
		" dbname=" + p.dbName +
		" user=" + p.dbUser +
		" password=" + p.dbPassword +
		" sslmode=" + p.dbSslMode

	db, err := sqlx.Connect(dbDriver, connectionString)
	if err != nil {
		logrus.Error(err.Error())
	}
	return db
}
