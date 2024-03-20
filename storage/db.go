package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func GetDB() (db *sql.DB) {

	if err := godotenv.Load("./config/.env"); err != nil {
		log.Print("No .env file found")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	sslmode := os.Getenv("SSLMODE")

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected to DB")

	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
