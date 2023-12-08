package store

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"net/http"
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

	return storeDB, nil
}

func (s *StoreDB) Create(originalURL, shortURL string) error {
	query := `
        INSERT INTO urls (short_id, original_url) 
        VALUES ($1, $2)
    `
	_, err := s.db.Exec(query, shortURL, originalURL)
	if err != nil {
		fmt.Println("error save URL: %v", err)
		return err
	}
	fmt.Println("URL save")
	return nil
}

func (s *StoreDB) Get(shortURL string) (string, error) {
	query := `
        SELECT original_url 
        FROM urls 
        WHERE short_id = $1
    `

	var originalURL string
	// Выполнение запроса к базе данных
	err := s.db.QueryRow(query, shortURL).Scan(&originalURL)
	if err != nil {
		// Обработка ошибки при выборке данных
		if err == sql.ErrNoRows {
			fmt.Println("short URL not fund", http.StatusNotFound)
			return "", err
		}
		fmt.Println(fmt.Sprintf("Ошибка при извлечении URL: %v", err), http.StatusInternalServerError)
		return "", err
	}

	return originalURL, err
}

func (s *StoreDB) PingStore() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := s.db.PingContext(ctx); err != nil {
		return fmt.Errorf("pinging db-store: %w", err)
	}
	return nil
}
