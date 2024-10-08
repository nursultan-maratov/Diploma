package main

import (
	_ "github.com/lib/pq"
	"github.com/nursultan-maratov/Diploma.git/internal/postgres"
	"log"
)

func main() {

	_, err := postgres.ConnectDefaultDataBase()
	if err != nil {
		log.Fatalf("Kaput db is not workings %v", err)
	}

}
