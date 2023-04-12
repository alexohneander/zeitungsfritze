package helper

import (
	"math/rand"
)

var pizzaIngredients = []string{
	"cheese",
	"pepperoni",
	"mushrooms",
	"sausage",
	"bacon",
	"pineapple",
	"onions",
	"olives",
	"spinach",
	"tomatoes",
	"ham",
	"chicken",
	"beef",
	"anchovies",
	"jalapenos",
	"garlic",
}

// use 3 slices from pizzaIngredients to generate domain names
func CreateDomain() (domain string) {
	for i := 0; i < 3; i++ {
		domain += pizzaIngredients[rand.Intn(len(pizzaIngredients))]
		if i < 2 {
			domain += "-"
		}
	}
	return domain
}
