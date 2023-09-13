package workshop

type Person struct {
	Name   string
	Wallet Wallet
}

type Wallet struct {
	Cash Money
}

type Money struct {
	Amount int
}

type Shop struct {
	Name    string
	Shelves []Shelf
}

type Shelf struct {
	Name     string
	Products []Product
}

type Product struct {
	Name     string
	Price    Money
	Quantity int
}

func HandlePersonPurchaseProducts(person *Person, shop *Shop, productName string, quantity int) bool {
	return false
}
