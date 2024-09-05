package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type PostgresqlRepository struct {
	DB *pgxpool.Pool
}

func NewPostgresqlRepository(ctx context.Context) (*PostgresqlRepository, error) {
	sqlStorage := &PostgresqlRepository{}

	err := sqlStorage.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return sqlStorage, nil
}

func (p *PostgresqlRepository) Connect(ctx context.Context) error {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	if user == "" || password == "" || name == "" || host == "" || port == "" {
		return fmt.Errorf("missing required environment variables")
	}

	connString := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + name

	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return err
	}

	p.DB = pool

	return nil
}

func (p *PostgresqlRepository) Close(ctx context.Context) {
	p.DB.Close()
}
