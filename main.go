package main

import "fmt"

// data contract
type Game struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
}

func main() {

	game := Game{
		ID:    1,
		Name:  "Game Testing",
		Genre: "Action",
	}

	fmt.Print(game)
	fmt.Println("AI Backend Go Started")
}
