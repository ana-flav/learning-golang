package data 

import (
	"database/sql"
	"fmt"
	"log"
	"os"

 	"github.com/joho/godotenv"
  
	_ "github.com/lib/pq"
  )
  

  func DataBase() (*sql.DB, error) {

	err := godotenv.Load(".env")
	if err != nil{
	 log.Fatalf("Error loading .env file: %s", err)
	}


	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))

	db, err := sql.Open("postgres", psqlInfo)

	return db, err
  }
	

