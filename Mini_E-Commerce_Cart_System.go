package main

import (
	"fmt"
	"log"
)

type Product struct {
	Name  string
	Price float64
}

type Products struct {
	List []Product
}

type CartItem struct {
	Product  Product
	Quantity int
}

type Cart struct {
	Items []CartItem
}

func (self *Products) fillProducts(name string, price float64) {
	if price <= 0 {
		log.Fatal("Value cannot be 0 or less")
	} else if name == "" {
		log.Fatal("Name cannot be empty")
	}
	product := Product{Name: name, Price: price}

	self.List = append(self.List, product)
}

func (self *Products) itemExists(name string) bool {
	for _, item := range self.List {
		if item.Name == name {
			return true
		}
	}
	return false
}

func (self *Cart) AddProduct(product *Product, quantity int, productList *Products) {
	if !productList.itemExists(product.Name) {
		log.Fatal("item does not exists ... ignoring")
	}
	if quantity == 0 {
		log.Fatal("quantity cannot be 0")
	}

	cardItem := CartItem{Product: *product, Quantity: quantity}
	self.Items = append(self.Items, cardItem)
	fmt.Printf("item added to Cart : %v (%v)\n", cardItem.Product.Name, cardItem.Quantity)
}

func (self *Cart) RemoveProduct(productName string) bool {
	for i, item := range self.Items {
		if item.Product.Name == productName {
			self.Items = append(self.Items[:i], self.Items[i+1:]...)
			return true
		} else {
			log.Fatal("Your cart does not contain that item")
			return false
		}
	}
	return false
}

func (self *Cart) GetTotalPrice(discount bool) float64 {
	// var sum float64
	sum := 0.0
	for _, item := range self.Items {
		sum += (item.Product.Price * float64(item.Quantity))
	}
	if discount {
		if sum >= 500.00 {
			fmt.Println("Total Price: ", sum)
			fmt.Println("-- [10 percent discount applied] --")
			sum -= (sum * 10.0 / 100)
		}
	}
	return sum
}

func (self *Cart) ListItems() {
	fmt.Println("Cart : ")
	for _, item := range self.Items {
		fmt.Printf(" - %v (%v)\n", item.Product.Name, item.Quantity)
	}
	fmt.Println()
}

func main() {
	var products Products
	var yourCart Cart

	products.fillProducts("Laptop", 1200.00)             // 0
	products.fillProducts("Smartphone", 800.00)          // 1
	products.fillProducts("Headphones", 150.00)          // 2
	products.fillProducts("Wireless Mouse", 40.00)       // 3
	products.fillProducts("Mechanical Keyboard", 110.00) // 4
	products.fillProducts("Smartwatch", 250.00)          // 5
	products.fillProducts("Backpack", 70.00)             // 6
	products.fillProducts("Monitor", 300.00)             // 7
	products.fillProducts("External Hard Drive", 90.00)  // 8
	products.fillProducts("USB-C Charger", 30.00)        // 9
	products.fillProducts("Gaming Chair", 400.00)        // 10
	products.fillProducts("Bluetooth Speaker", 120.00)   // 11

	yourCart.AddProduct(&products.List[6], 1, &products)
	yourCart.AddProduct(&products.List[2], 1, &products)
	yourCart.AddProduct(&products.List[4], 1, &products)
	yourCart.AddProduct(&products.List[5], 1, &products)

	yourCart.ListItems()

	fmt.Printf("Total Price : %v\n", yourCart.GetTotalPrice(true))

}
