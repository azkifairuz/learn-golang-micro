package main

import (
	"auth/data"
	"database/sql"
	"log"
)


const webPort = "80"

type Config struct {
	DB *sql.DB
	Models data.Models
}

func main()  {
	log.Println("start authentication service")
}