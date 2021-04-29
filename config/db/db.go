package db

import (
	"database/sql"
	"fmt"
	"go-rest-api/config/dbconst"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func OpenConnectionDb() *sql.DB {
	db, err := sql.Open("postgres", configUri())
	if err != nil {
		panic(err.Error())
	}

	return db
}

func configUri() string {
	envs, err := godotenv.Read()
	if err != nil {
		panic(err.Error())
	}
	uriConnection := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", envs[dbconst.Dbuser],
		envs[dbconst.Dbname], envs[dbconst.Dbpass], envs[dbconst.Dbhost])

	return uriConnection
}
