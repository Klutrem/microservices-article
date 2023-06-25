package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database interface {
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
}
type Client interface {
	Ping(context.Context) error
	Connect(context.Context, string) error
	Disconnect()
	Database() Database
}
type PostgresClient struct {
	pool *pgxpool.Pool
}

func NewClient() *PostgresClient {
	return &PostgresClient{}
}
func (p *PostgresClient) Ping(ctx context.Context) error {
	return p.pool.Ping(ctx)
}
func (p *PostgresClient) Connect(ctx context.Context, connectionString string) error {
	config, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return err
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return err
	}
	p.pool = pool
	return nil
}
func (p *PostgresClient) Disconnect() {
	if p.pool != nil {
		p.pool.Close()
	}
}
func (p *PostgresClient) Database() Database {
	return p
}
func (p *PostgresClient) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return p.pool.QueryRow(ctx, sql, args...)
}
func (p *PostgresClient) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return p.pool.Query(ctx, sql, args...)
}
func (p *PostgresClient) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	return p.pool.Exec(ctx, sql, args...)
}
