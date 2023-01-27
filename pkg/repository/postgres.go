package repository

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func AddVidesToDatabase(db *pgxpool.Pool) error {
	for i := 1; i < 5; i++ {
		filename := fmt.Sprintf("mock_videos_%d.sql", i)
		readFile, err := os.Open(filepath.Join("mockdata", filename))
		if err != nil {
			return err
		}

		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			if _, err = db.Exec(context.Background(), fileScanner.Text()); err != nil {
				return err
			}
		}
	}
	return nil
}
