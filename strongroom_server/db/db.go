package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func CreateAccount(username string, display_name string, auth_salt, muk_salt, auth_verifier []byte) {
	dbinfo := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable",
		"db", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_DB"))
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("# Inserting values")

	var lastInsertId int
	sql := "INSERT INTO accounts(username, display_name, auth_salt, muk_salt, auth_verifier) " +
		"VALUES($1,$2,$3,$4,$5) returning id;"
	err = db.QueryRow(sql, username, display_name, auth_salt, muk_salt, auth_verifier).Scan(&lastInsertId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("last inserted id =", lastInsertId)
}
