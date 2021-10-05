package config

import (
	"database/sql"
	"encoding/json"
	"log"
)

func CreateConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:alen@localhost/go?sslmode=require")
	if err != nil {
		log.Fatal(err)
	}

	// validasi konfigurasi dengan ping ke database server
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("connected successfully!")
	return db
}

type NullString struct {
	sql.NullString
}

func (s NullString) MarshalJson() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}
func (s *NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		s.String, s.Valid = "", false
		return nil
	}
	s.String, s.Valid = string(data), true
	return nil
}
