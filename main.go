package main

import (
	"encoding/json"
	"fmt"
	"goj/migration"
	"goj/repo"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	connect()
	defer disconnect()

	migration.Seed(database)

	http.HandleFunc("/albums", func(w http.ResponseWriter, r *http.Request) {
		albums, err := repo.GetAlbums(database)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}

		json.NewEncoder(w).Encode(albums)
	})

	http.ListenAndServe(":3000", nil)

}
