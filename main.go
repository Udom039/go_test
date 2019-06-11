package main

import "fmt"
import "encoding/json"

type customer struct {
	Firstname string
	Lastname  string
	Code      int
	Phone     string
}

func main() {
	customerLists := []customer{}
	customer1 := customer{
		Firstname: "Chaiyarin",
		Lastname:  "Niamsuwan",
		Code:      111990,
		Phone:     "085661234",
	}
	customer2 := customer{
		Firstname: "Atikom",
		Lastname:  "Sombutjalearn",
		Code:      111991,
		Phone:     "085664321",
	}
	customer3 := customer{
		Firstname: "Kritsana",
		Lastname:  "Punyaphon",
		Code:      111992,
		Phone:     "085662344",
	}
	customerLists = append(customerLists, customer1)
	customerLists = append(customerLists, customer2)
	customerLists = append(customerLists, customer3)
	customerListsJson, err := json.Marshal(customerLists) // Marshal... Array to JSON

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(customerListsJson))
}

//test git////
