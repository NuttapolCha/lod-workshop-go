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
	for _, shelf := range shop.Shelves {
		for i, product := range shelf.Products {
			if product.Name == productName {
				if person.Wallet.Cash.Amount >= product.Price.Amount*quantity && product.Quantity >= quantity {
					person.Wallet.Cash.Amount -= product.Price.Amount * quantity
					shelf.Products[i].Quantity -= quantity
					return true
				}
			}
		}
	}
	return false
}
