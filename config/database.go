package config

import(
	"database/sql"
	"fmt"
	"log"
	"os"
	
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB () *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error load env files")
	}
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_TIMEZONE"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("error connect dtabase: ", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("database gabisa diakses: ", err)
	}
	log.Println("database uda connect")
	return db
}