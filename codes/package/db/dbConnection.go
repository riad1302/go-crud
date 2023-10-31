package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	config "golang-crud/configs"

	_ "github.com/go-sql-driver/mysql"
)

func dsn() string {
	connection := config.GetConnectionString()
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", connection.Username, connection.Password, connection.Hostname, connection.Dbname)
}

func DbConnection() (*sql.DB, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	db, err := sql.Open("mysql", dsn())
	if err != nil {
		log.Printf("Error %s when opening db\n", err)
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s when pinging database\n", err)
		return nil, err
	}

	log.Printf("Connected to DB %s successfully\n", "dbname")

	return db, nil
}
