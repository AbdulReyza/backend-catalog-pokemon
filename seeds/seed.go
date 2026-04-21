package main

import (
	"log"

	"github.com/AbdulReyza/backend-catalog-pokemon.git/config"
	"github.com/AbdulReyza/backend-catalog-pokemon.git/models"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.InitDatabase()
	products := []models.Product{
		{
			Name:        "Pikachu",
			Price:       50000,
			Category:    "Rare",
			Stock:       10,
			Description: "Kartu Pikachu edisi holographic dengan serangan listrik",
			ImageURL:    "https://i.ibb.co.com/0pcntPhh/65d08057328b513c7148c65e-pokemon-first-partner-pikachu-oversize.jpg",
		},
		{
			Name:        "Charizard",
			Price:       150000,
			Category:    "Rare",
			Stock:       5,
			Description: "Kartu Charizard langka dengan efek api kuat",
			ImageURL:    "https://i.ibb.co.com/Y4YCCrLn/il-570x-N-4140836231-cpmu.webp",
		},
		{
			Name:        "Mewtwo",
			Price:       120000,
			Category:    "Rare",
			Stock:       7,
			Description: "Kartu Mewtwo rare dengan kekuatan psychic",
			ImageURL:    "https://i.ibb.co.com/mVwRXq4C/id00014444.png",
		},
		{
			Name:        "Bulbasaur",
			Price:       15000,
			Category:    "Common",
			Stock:       25,
			Description: "Kartu Bulbasaur basic tipe rumput",
			ImageURL:    "https://i.ibb.co.com/q3Sk7Jdp/PGO-EN-1.png",
		},
		{
			Name:        "Squirtle",
			Price:       15000,
			Category:    "Common",
			Stock:       30,
			Description: "Kartu Squirtle basic tipe air",
			ImageURL:    "https://m.media-amazon.com/images/I/51qK-vD4z5L._AC_UF894,1000_QL80_.jpg",
		},
		{
			Name:        "Charmander",
			Price:       15000,
			Category:    "Common",
			Stock:       28,
			Description: "Kartu Charmander basic tipe api",
			ImageURL:    "https://m.media-amazon.com/images/I/81XGIXiel2L._AC_UF894,1000_QL80_.jpg",
		},
		{
			Name:        "Pidgey",
			Price:       10000,
			Category:    "Common",
			Stock:       35,
			Description: "Kartu Pidgey common tipe terbang",
			ImageURL:    "https://i.ibb.co.com/RpBwB13M/id00008932.png",
		},
		{
			Name:        "Rattata",
			Price:       8000,
			Category:    "Common",
			Stock:       40,
			Description: "Kartu Rattata common tipe normal",
			ImageURL:    "https://i.ibb.co.com/tTwz9ZTP/id00008430-1.png",
		},
	}
	for _, p := range products {
		config.DB.Create(&p)
	}
	log.Printf("Seed berhasil: %d produk ditambahkan", len(products))
}
