package postgresql

import (
	"context"
	"database/sql"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Database interface {
	// QueryRowx queries the database and returns an *sqlx.Row.
	// Any placeholder parameters are replaced with supplied args.
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	// Queryx queries the database and returns an *sqlx.Rows.
	// Any placeholder parameters are replaced with supplied args.
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	// Select using this DB.
	// Any placeholder parameters are replaced with supplied args.
	Select(dest interface{}, query string, args ...interface{}) error
	//MustExec (panic) runs MustExec using this database.
	//Any placeholder parameters are replaced with supplied args.
	MustExec(query string, args ...interface{}) sql.Result
	//Exec executes a query without returning any rows.
	//The args are for any placeholder parameters in the query.
	Exec(query string, args ...any) (sql.Result, error)
	//NamedExec using this DB.
	//Any named placeholder parameters are replaced with fields from arg
	NamedExec(query string, arg interface{}) (sql.Result, error)
	//	Get using this DB.
	// Any placeholder parameters are replaced with supplied args.
	// An error is returned if the result set is empty.
	Get(dest interface{}, query string, args ...interface{}) error
}

type Client interface {
	Ping(context.Context) error
	Disconnect()
	Database() Database
}

type PostgresClient struct {
	pool *pgxpool.Pool
	db   *sqlx.DB
}

func NewClient(connectionString string) (Client, error) {
	p := PostgresClient{}
	config, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return p, err
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return p, err
	}
	p.pool = pool
	nativeDB := stdlib.OpenDB(*config.ConnConfig)

	db := sqlx.NewDb(nativeDB, "pgx")
	p.db = db

	migration, err := os.ReadFile("pkg/postgresql/migration/users.sql")
	if err != nil {
		return p, err
	}
	db.Exec(string(migration))
	return p, nil
}

func (p PostgresClient) Ping(ctx context.Context) error {
	err := p.pool.Ping(ctx)
	if err != nil {
		return err
	}
	return p.db.Ping()
}

func (p PostgresClient) Disconnect() {
	if p.pool != nil {
		p.pool.Close()
		p.db.Close()
	}
}

func (p PostgresClient) Database() Database {
	return p.db
}
