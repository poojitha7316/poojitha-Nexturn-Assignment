package main

import (
	"fmt"
	"sort"
	"strconv"
)

// Product struct represents a product in the inventory
type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

// Inventory is a slice that holds all products in the inventory
var Inventory []Product

// AddProduct function adds a new product to the inventory
func AddProduct(id int, name string, priceInput string, stock int) error {
	// Convert priceInput to float64
	price, err := strconv.ParseFloat(priceInput, 64)
	if err != nil {
		return fmt.Errorf("invalid price format: %v", err)
	}

	// Create new product
	newProduct := Product{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	}

	// Add the product to inventory
	Inventory = append(Inventory, newProduct)
	return nil
}

// UpdateStock function updates the stock of a product
func UpdateStock(id int, newStock int) error {
	if newStock < 0 {
		return fmt.Errorf("stock cannot be negative")
	}

	for i, product := range Inventory {
		if product.ID == id {
			Inventory[i].Stock = newStock
			return nil
		}
		fmt.Println("Inventory after updation of stock is:")

	}
	return fmt.Errorf("product with ID %d not found", id)
}

// SearchProduct function allows searching for a product by ID or Name
func SearchProduct(query string) (*Product, error) {
	for _, product := range Inventory {
		// Search by name
		if product.Name == query {
			return &product, nil
		}

		// Search by ID (convert query to integer)
		if id, err := strconv.Atoi(query); err == nil && product.ID == id {
			return &product, nil
		}
	}
	return nil, fmt.Errorf("product not found")
}

// DisplayInventory function prints the inventory table
func DisplayInventory() {
	fmt.Printf("%-5s %-20s %-10s %-5s\n", "ID", "Name", "Price", "Stock")
	fmt.Println("---------------------------------------------------")
	for _, product := range Inventory {
		fmt.Printf("%-5d %-20s %-10.2f %-5d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

// SortInventoryByPrice function sorts products by price in ascending order
func SortInventoryByPrice() {
	sort.SliceStable(Inventory, func(i, j int) bool {
		return Inventory[i].Price < Inventory[j].Price
	})
	fmt.Println("Inventory after sorting the stock by price:")

}

// SortInventoryByStock function sorts products by stock in ascending order
func SortInventoryByStock() {
	sort.SliceStable(Inventory, func(i, j int) bool {
		return Inventory[i].Stock < Inventory[j].Stock
	})
	fmt.Println("Inventory after sorting the stock by stock:")

}

func main() {
	// Sample products for testing
	if err := AddProduct(1, "mobile", "1000.00", 100); err != nil {
		fmt.Println("Error:", err)
	}
	if err := AddProduct(2, "laptop", "500.00", 50); err != nil {
		fmt.Println("Error:", err)
	}
	if err := AddProduct(3, "camera", "900.00", 75); err != nil {
		fmt.Println("Error:", err)
	}

	// Display the initial inventory
	fmt.Println("Initial Inventory:")
	DisplayInventory()

	// Search for a product
	fmt.Println("\nSearching for product 'camera':")
	product, err := SearchProduct("camera")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Found product: %v\n", product)
	}

	// Update stock for product ID 2 (laptop)
	fmt.Println("\nUpdating stock for laptop (ID 2) to :90")
	if err := UpdateStock(2, 90); err != nil {
		fmt.Println("Error:", err)
	} else {
		DisplayInventory()
	}

	// Sort by price and display the inventory
	SortInventoryByPrice()
	DisplayInventory()

	// Sort by stock and display the inventory
	SortInventoryByStock()
	DisplayInventory()
}


