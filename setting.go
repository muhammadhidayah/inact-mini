package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Setting struct{}

func NewSetting() *Setting {
	return &Setting{}
}

func (s *Setting) Connect() (*sql.DB, error) {
	connString := s.configConStr()
	db, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (s *Setting) configConStr() string {
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "docker"
	dbName := "inact_mini"

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	return connStr
}
