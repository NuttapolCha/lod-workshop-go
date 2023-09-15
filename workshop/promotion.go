package workshop

type Promotion struct {
	Discount int
}

func HandlePersonPurchaseProductsWithPromotion(person *Person, shop *Shop, promotion Promotion, productName string, quantity int) bool {
	for _, shelf := range shop.Shelves {
		for i, product := range shelf.Products {
			if product.Name == productName {
				if person.Wallet.Cash.Amount >= product.Price.Amount*quantity-promotion.Discount && product.Quantity >= quantity {
					person.Wallet.Cash.Amount -= product.Price.Amount*quantity - promotion.Discount
					shelf.Products[i].Quantity -= quantity
					return true
				}
			}
		}
	}
	return false
}
