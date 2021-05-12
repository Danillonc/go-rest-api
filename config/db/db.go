package db

import (
	"database/sql"
	"fmt"
	"go-rest-api/config/db/dbconst"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//Function responsib le to open connection with a database
func OpenConnectionDb() *sql.DB {
	db, err := sql.Open("postgres", configUri())
	if err != nil {
		panic(err.Error())
	}
	return db
}

func configUri() string {
	envs, err := godotenv.Read("config/env/.env")
	if err != nil {
		panic(err.Error())
	}
	uriConnection := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", envs[dbconst.Dbuser],
		envs[dbconst.Dbname], envs[dbconst.Dbpass], envs[dbconst.Dbhost])

	return uriConnection
}
