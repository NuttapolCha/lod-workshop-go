package workshop

type Person struct {
	Name   string
	Wallet Wallet
}

type Wallet struct {
	Cash Money
}

func (w Wallet) Balance() int {
	return w.Cash.GetAmount()
}

func (w Wallet) CanSpendBy(amount int) bool {
	return w.Balance() >= amount
}

func (w *Wallet) Spend(amount int) {
	w.Cash.Add(-1 * amount)
}

type Money struct {
	Amount int
}

func (m Money) GetAmount() int {
	return m.Amount
}

func (m *Money) Add(amount int) {
	m.Amount += amount
}

type Shop struct {
	Name    string
	Shelves []Shelf
}

func (s Shop) CalculatePricing(productName string, quantity int) (int, bool) {
	for _, shelf := range s.Shelves {
		if product := shelf.FindProductByName(productName); product != nil {
			return product.TotalPrice(quantity), true
		}
	}
	return 0, false
}

func (s *Shop) MakePurchase(productName string, quantity int) bool {
	for _, shelf := range s.Shelves {
		if shelf.HasEnoughProduct(productName, quantity) {
			shelf.ReduceProductQuantity(productName, quantity)
			return true
		}
	}
	return false
}

type Shelf struct {
	Name     string
	Products []Product
}

func (s Shelf) FindProductByName(name string) *Product {
	for _, p := range s.Products {
		if p.Name == name {
			return &p
		}
	}
	return nil
}

func (s Shelf) HasEnoughProduct(productName string, quantity int) bool {
	product := s.FindProductByName(productName)
	if product == nil {
		return false
	}

	return product.HaveEnough(quantity)
}

func (s *Shelf) ReduceProductQuantity(productName string, quantity int) {
	for i, p := range s.Products {
		if p.Name == productName {
			s.Products[i].ReduceQuantity(quantity)
			// p.ReduceQuantity(quantity)
		}
	}
}

type Product struct {
	Name     string
	Price    Money
	Quantity int
}

func (p Product) TotalPrice(quantity int) int {
	return p.Price.GetAmount() * quantity
}

func (p Product) HaveEnough(quantity int) bool {
	return p.Quantity >= quantity
}

func (p *Product) ReduceQuantity(quantity int) {
	p.Quantity -= quantity
}

func HandlePersonPurchaseProducts(person *Person, shop *Shop, productName string, quantity int) bool {
	buyingPrice, inStock := shop.CalculatePricing(productName, quantity)
	if !inStock {
		return false
	}

	if !person.Wallet.CanSpendBy(buyingPrice) {
		return false
	}

	isPurchased := shop.MakePurchase(productName, quantity)
	if !isPurchased {
		return false
	}

	person.Wallet.Spend(buyingPrice)
	return true
}
