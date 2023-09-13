package workshop_test

import (
	"testing"

	"github.com/NuttapolCha/lod-workshop-go/workshop"
	"github.com/stretchr/testify/assert"
)

func NewTestingPerson() *workshop.Person {
	return &workshop.Person{
		Name: "Youth",
		Wallet: workshop.Wallet{
			Cash: workshop.Money{
				Amount: 100,
			},
		},
	}
}

func NewTestingShop() *workshop.Shop {
	return &workshop.Shop{
		Name: "7-11",
		Shelves: []workshop.Shelf{
			{
				Name: "food",
				Products: []workshop.Product{
					{
						Name: "bread",
						Price: workshop.Money{
							Amount: 50,
						},
						Quantity: 2,
					},
					{
						Name: "egg",
						Price: workshop.Money{
							Amount: 20,
						},
						Quantity: 2,
					},
				},
			},
			{
				Name: "drink",
				Products: []workshop.Product{
					{
						Name: "water",
						Price: workshop.Money{
							Amount: 10,
						},
						Quantity: 2,
					},
					{
						Name: "milk",
						Price: workshop.Money{
							Amount: 20,
						},
						Quantity: 2,
					},
					{
						Name: "whey protein",
						Price: workshop.Money{
							Amount: 100,
						},
						Quantity: 2,
					},
				},
			},
		},
	}
}

func TestHandlePersonPurchaseProduct(t *testing.T) {
	t.Run("person with enough money", func(t *testing.T) {
		initialPerson := NewTestingPerson()
		initialShop := NewTestingShop()

		buyingProduct := "bread"
		expectedPersonAfterBuying := &workshop.Person{
			Name: "Youth",
			Wallet: workshop.Wallet{
				Cash: workshop.Money{
					Amount: 50,
				},
			},
		}
		expectedShopAfterBuying := &workshop.Shop{
			Name: "7-11",
			Shelves: []workshop.Shelf{
				{
					Name: "food",
					Products: []workshop.Product{
						{
							Name: "bread",
							Price: workshop.Money{
								Amount: 50,
							},
							Quantity: 1,
						},
						{
							Name: "egg",
							Price: workshop.Money{
								Amount: 20,
							},
							Quantity: 2,
						},
					},
				},
				{
					Name: "drink",
					Products: []workshop.Product{
						{
							Name: "water",
							Price: workshop.Money{
								Amount: 10,
							},
							Quantity: 2,
						},
						{
							Name: "milk",
							Price: workshop.Money{
								Amount: 20,
							},
							Quantity: 2,
						},
						{
							Name: "whey protein",
							Price: workshop.Money{
								Amount: 100,
							},
							Quantity: 2,
						},
					},
				},
			},
		}

		workshop.HandlePersonPurchaseProduct(initialPerson, initialShop, buyingProduct)
		assert.Equal(t, expectedPersonAfterBuying, initialPerson)
		assert.Equal(t, expectedShopAfterBuying, initialShop)
	})
}
