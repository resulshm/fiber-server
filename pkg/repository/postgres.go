package repository

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
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

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AddVidesToDatabase(db *sqlx.DB) error {
	for i := 1; i < 5; i++ {
		filename := fmt.Sprintf("mock_videos_%d.sql", i)
		readFile, err := os.Open(filepath.Join("mockdata", filename))
		if err != nil {
			return err
		}

		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			if _, err = db.Exec(fileScanner.Text()); err != nil {
				return err
			}
		}
	}
	return nil
}
