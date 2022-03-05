package main

import ("fmt"
		"ddd-workshop/internal/pkg/domain"
	)

func main() {

	fmt.Println("Hello DDD")
	cart := domain.Cart{
	}
	
	anotherCart := domain.Cart{
	}

	cart.AddItemToCart(domain.Item{
		domain.Product{"ipad", "123"},
		2})
	
	cart.AddItemToCart(domain.Item{
		domain.Product{"GM cricket bat","233"},
		2})

	anotherCart.AddItemToCart(domain.Item{
		domain.Product{"ipad", "123"},
		2})
		
	removedFromCart := cart.RemoveFromCart(domain.Item{
		domain.Product{"ipad", "123"},
		2})

	removedItem := domain.RemovedItems{}

	removedItem.Add(removedFromCart)
	
	fmt.Println(cart.Items)
	fmt.Println(removedItem.Name)
	// fmt.Println(cart.IsIdentifcalTo(anotherCart))

}