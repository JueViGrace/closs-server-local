package data

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type Storage interface {
	AuthStore() AuthStore
	SessionStorage() SessionStore
	OrderStorage() OrderStore
	ProductStore() ProductStore
	UserStore() UserStore

	Health() []map[string]string
	Close() error
}

type storage struct {
	myStore    *mySQLStore
	cacheStore *cacheStore
	ctx        context.Context
}

var ctx = context.Background()

func NewStorage() Storage {
	return &storage{
		myStore:    newMySQLStorage(),
		cacheStore: newCacheStorage(),
		ctx:        ctx,
	}
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *storage) Health() []map[string]string {
	healthS := make([]map[string]string, 0)

	healthS = append(healthS, s.myStore.health())
	healthS = append(healthS, s.cacheStore.health())

	return healthS
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *storage) Close() error {
	log.Printf("Disconnected from database: %s", dbName)
	log.Printf("Disconnected from database: %s", dbUrl)

	if err := s.myStore.close(); err != nil {
		return err
	}

	if err := s.cacheStore.close(); err != nil {
		return err
	}

	return nil
}
