package main

import "fmt"

type OrderItem struct {
	Product string `json:"product"`
	Price   int    `json:"price"`
}
type Order struct {
	OrderID int         `json:"order_id"`
	Items   []OrderItem `json:"items"`
}

type User struct {
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
}

func main() {

	order := Order{
		OrderID: 123,
		Items: []OrderItem{
			{Product: "Laptop", Price: 1000},
			{Product: "Mouse", Price: 50},
		},
	}

	User := User{
		Name:  "Galuh",
		Roles: []string{"admin", "editor"},
	}

	fmt.Println(order)
	fmt.Println(User)
	fmt.Println("AI Backend Go Started")
}
