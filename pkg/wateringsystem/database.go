package wateringsystem

import "github.com/jackc/pgx"

func NewDatabase() (*pgx.ConnPool, error) {
	config := pgx.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		Database: "watersystem",
		User:     "pomelo",
		Password: "pomelo",
	}

	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{ConnConfig: config})
	if err != nil {
		return nil, err
	}

	return pool, err
}
