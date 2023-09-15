package workshop

type Person struct {
	Name   string
	Wallet Wallet
}

func (p *Person) Spend(amount int) bool {
	if p.Wallet.CanSpendBy(amount) {
		p.Wallet.DeductMoney(amount)
		return true
	}
	return false
}

type Wallet struct {
	Cash Money
}

func (w Wallet) CanSpendBy(amount int) bool {
	return w.Cash.CanSpendBy(amount)
}

func (w *Wallet) DeductMoney(amount int) {
	w.Cash.Deduct(amount)
}

type Money struct {
	Amount int
}

func (m Money) CanSpendBy(amount int) bool {
	return m.Amount >= amount
}

func (m *Money) Deduct(amount int) {
	m.Amount -= amount
}

type Shop struct {
	Name    string
	Shelves []Shelf
}

func (s *Shop) FindProductFromAnyShelf(name string) *Product {
	for _, shelf := range s.Shelves {
		product := shelf.FindProduct(name)
		if product != nil {
			return product
		}
	}
	return nil
}

type Shelf struct {
	Name     string
	Products []Product
}

func (s *Shelf) FindProduct(name string) *Product {
	for i, product := range s.Products {
		if product.Name == name {
			return &s.Products[i]
		}
	}
	return nil
}

type Product struct {
	Name     string
	Price    Money
	Quantity int
}

func (p Product) CheckPrice(quantity int) int {
	return p.Price.Amount * quantity
}

func (p Product) HaveEnough(quantity int) bool {
	return p.Quantity >= quantity
}

func (p *Product) DeductQuantity(quantity int) {
	p.Quantity -= quantity
}

func HandlePersonPurchaseProducts(person *Person, shop *Shop, productName string, quantity int) bool {
	product := shop.FindProductFromAnyShelf(productName)
	if product == nil {
		return false
	}

	if !product.HaveEnough(quantity) {
		return false
	}

	toSpend := product.CheckPrice(quantity)
	isSpent := person.Spend(toSpend)
	if !isSpent {
		return false
	}

	product.DeductQuantity(quantity)

	return true
}
