package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/mathyourlife/rbac-handlers/pkg/backend/postgres"
	"github.com/mathyourlife/rbac-handlers/pkg/rbmw"
)

var logger = log.New(os.Stderr, "[auth-server] ", log.LstdFlags|log.Lshortfile)

type config struct {
	PG *postgres.Config
}

func parseFlags() (*config, error) {
	c := &config{
		PG: &postgres.Config{},
	}

	flag.StringVar(&c.PG.Host, "db.host", "localhost", "database hostname")
	flag.IntVar(&c.PG.Port, "db.port", 5432, "database port")
	flag.StringVar(&c.PG.Username, "db.username", "", "username")
	flag.StringVar(&c.PG.Password, "db.password", "", "password")
	flag.StringVar(&c.PG.DBName, "db.dbname", "", "database name")
	flag.BoolVar(&c.PG.SSLEnabled, "db.ssl", false, "enable ssl")

	flag.Parse()

	return c, nil
}

func main() {
	logger.Println("starting api server")

	c, err := parseFlags()
	if err != nil {
		logger.Fatal(err)
	}

	db, err := postgres.NewDB(c.PG)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	var h *rbmw.Handler
	if h, err = rbmw.NewHandler(&rbmw.HandlerConfig{Backend: db}); err != nil {
		panic(err)
	}

	h.AddHandlers(mux, "/auth")

	log.Fatal(http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", mux))
}
