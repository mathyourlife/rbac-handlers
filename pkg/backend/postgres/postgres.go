package postgres

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Config struct {
	Host       string
	Port       int
	Username   string
	Password   string
	DBName     string
	SSLEnabled bool
}

func (c *Config) dsn() string {
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		c.Host, c.Port, c.DBName, c.Username, c.Password)
}

func NewDB(config *Config) (*DB, error) {
	gdb, err := gorm.Open("postgres", config.dsn())
	if err != nil {
		return nil, err
	}
	db := &DB{
		gdb.Debug(),
	}
	return db, nil
}

type DB struct {
	*gorm.DB
}
