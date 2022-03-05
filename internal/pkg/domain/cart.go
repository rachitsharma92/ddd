package domain

import "github.com/thoas/go-funk" 

type Cart struct{
	Items []Item
}

func (cart *Cart) AddItemToCart (item Item) {
	cart.Items = append(cart.Items,item)
}

func (cart Cart) IsIdentifcalTo (anotherCart Cart) bool {
	return  funk.Equal(cart, anotherCart) 
}

func (cart *Cart) RemoveFromCart (item Item) string {
	itemToRemove := item.Product.Name

	index := cart.getIndexOfItem(itemToRemove)

	if (index != -1){
		cart.Items = append(cart.Items[:index], cart.Items[index+1:]...)
		return itemToRemove
	}
	return ""
}

func (cart *Cart)getIndexOfItem(item string) int{
	for index, element := range cart.Items{
		if element.Product.Name == item{
			return index
		}
	}
	return -1
}