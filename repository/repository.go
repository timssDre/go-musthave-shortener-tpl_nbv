package repository

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"time"
)

type StoreDB struct {
	db *sql.DB
}

func InitDatabase(DatabasePath string) (*StoreDB, error) {
	db, err := sql.Open("pgx", DatabasePath)
	if err != nil {
		return nil, fmt.Errorf("error opening db: %w", err)
	}

	storeDB := new(StoreDB)
	storeDB.db = db

	if DatabasePath != "" {
		err = createTable(db)
		if err != nil {
			return nil, fmt.Errorf("error creae table db: %w", err)
		}
	}

	return storeDB, nil
}

func (s *StoreDB) Create(originalURL, shortURL, UserID string) error {
	query := `
        INSERT INTO urls (short_id, original_url, userID) 
        VALUES ($1, $2, $3)
    `
	_, err := s.db.Exec(query, shortURL, originalURL, UserID)
	if err != nil {
		return err
	}
	return nil
}

func createTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS urls (
		id SERIAL PRIMARY KEY,
		short_id VARCHAR(256) NOT NULL UNIQUE,
		original_url TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    	userID VARCHAR(360)
	);
	DO $$ 
	BEGIN 
   	 IF NOT EXISTS (SELECT 1 FROM pg_indexes WHERE tablename = 'urls' AND indexname = 'idx_original_url') THEN
        CREATE UNIQUE INDEX idx_original_url ON urls(original_url);
    END IF;
	END $$;`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (s *StoreDB) GetFull(userID string, BaseURL string) ([]map[string]string, error) {
	query := `SELECT short_id, original_url FROM urls WHERE userID = $1`
	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get links: %w", err)
	}
	defer rows.Close()

	urls := make([]map[string]string, 0)
	for rows.Next() {
		var shortID, originalURL string
		if err = rows.Scan(&shortID, &originalURL); err != nil {
			return nil, err
		}
		shortURL := fmt.Sprintf("%s/%s", BaseURL, shortID)
		urlMap := map[string]string{"short_id": shortURL, "original_url": originalURL}
		urls = append(urls, urlMap)
	}

	return urls, nil
}

func (s *StoreDB) Get(shortURL string, originalURL string) (string, error) {
	field1 := "original_url"
	field2 := "short_id"
	field := shortURL
	if shortURL == "" {
		field2 = "original_url"
		field1 = "short_id"
		field = originalURL
	}

	query := fmt.Sprintf(`
        SELECT %s 
        FROM urls 
        WHERE %s = $1
    `, field1, field2)

	var answer string
	err := s.db.QueryRow(query, field).Scan(&answer)
	if err != nil {
		return "", err
	}

	return answer, err
}

func (s *StoreDB) PingStore() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := s.db.PingContext(ctx); err != nil {
		return fmt.Errorf("pinging db-store: %w", err)
	}
	return nil
}
