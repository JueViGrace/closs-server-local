package data

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type Storage struct {
	MyStore    *MySQLStore
	CacheStore CacheStore
}

var ctx = context.Background()

func NewStorage() *Storage {
	return &Storage{
		MyStore:    newMySQLStorage(),
		CacheStore: newCacheStorage(),
	}
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *Storage) Health() []map[string]string {
	healthS := make([]map[string]string, 0)

	healthS = append(healthS, s.MyStore.health())
	healthS = append(healthS, s.CacheStore.health())

	return healthS
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *Storage) Close() error {
	log.Printf("Disconnected from database: %s", dbName)
	log.Printf("Disconnected from database: %s", dbUrl)

	if err := s.MyStore.close(); err != nil {
		return err
	}

	if err := s.CacheStore.close(); err != nil {
		return err
	}

	return nil
}
