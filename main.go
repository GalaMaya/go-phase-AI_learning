package main

import (
	"fmt"
	"strings"
)

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

func ValidateUser(u User) error {
	if u.Name == "" {
		return fmt.Errorf("name is required")
	}

	if len(u.Roles) == 0 {
		return fmt.Errorf("at least one role is required")
	}

	return nil
}

func ValidateOrder(o Order) error {
	if o.OrderID == 0 {
		return fmt.Errorf("order_id is required")
	}

	if len(o.Items) == 0 {
		return fmt.Errorf("order must have at least one item")
	}

	for i, item := range o.Items {
		if item.Product == "" {
			return fmt.Errorf("item[%d]: product is required", i)
		}

		if item.Price <= 0 {
			return fmt.Errorf("item[%d]: price must be greater than zero", i)
		}
	}

	return nil
}

func ValidationCreateUser(u User) error {
	if u.Name == "" {
		return fmt.Errorf("name is required")
	}

	if len(u.Roles) == 0 {
		return fmt.Errorf("at least one role is required")
	}
	return nil
}

func NormalizeUser(u *User) {
	for i, role := range u.Roles {
		u.Roles[i] = strings.TrimSpace(role)
	}
}

func CreateUser(u *User) error {

	if err := ValidationCreateUser(*u); err != nil {
		return fmt.Errorf("Create User failed : %w", err)
	}

	NormalizeUser(u)

	return nil
}

func main() {

	order := Order{
		OrderID: 123,
		Items: []OrderItem{
			{Product: "Laptop", Price: 1000},
			{Product: "Mouse", Price: 50},
		},
	}

	user := User{
		Name:  "Galuh",
		Roles: []string{"admin", "editor"},
	}

	if err := ValidateUser(user); err != nil {
		fmt.Println("User validation failed:", err)
		return
	}

	if err := ValidateOrder(order); err != nil {
		fmt.Println("Order validation failed:", err)
		return
	}

	fmt.Println(order)
	fmt.Println(user)
	fmt.Println("AI Backend Go Started")
}
