# LoD Mini Workshop Golang

## Introduction

This is a mini workshop for learning the LoD (Law of Demeter) in Golang.
You will learn by doing the exercises in this workshop.

## Setup

1. Clone this repository

   ```sh
   git clone https://github.com/NuttapolCha/lod-workshop-go.git
   ```

2. Checkout to main branch

    ```sh
    git checkout main
    ```

3. You will coding at `./workshop/workshop.go`

## Getting Started

You are given pre-defined structs. Your job is to implement the *HandlePersonPurchaseProducts* function that handle the `person` purchasing
any `quantity` of `productName` within a `shop`.

The person can purchase products if

1. The `person` has enough money, i.e. person's money after purchasing product must be greater than or equal to zero.
2. The `productName` is exist within the shop and has enough quantity.

After making purchase, the person's money and shop's product quantity must be updated.
The function return true if the person successfully purchase the products, otherwise return false.

### Constraints

1. You cannot modify the fields of the structs.
2. In case `productName` exists within the shop, it is guaranteed to be unique within a shelf.

### Submission

You can run the test script to check your implementation.

```sh
make test
```

### What's next?

Share your code with your team, discussing on maintainability, readability. Thinking of scenarios where new requirements are added.

> Is your code easy to change? Will your new code duplicate existing code? How can you improve your code?

### New Requirements

Checkout to `with-promotion` branch. You are given a new requirement that a person can purchase a quantity of products with promotion.
The promotion contains only discount field which is the fixed discount for purchasing products.

```sh
git checkout with-promotion
```

Try to implement the *HandlePersonPurchaseProductsWithPromotion* then test if your code satisfy the requirements.

```sh
make test
```

## Appendix

### LoD

The Law of Demeter (LoD) or principle of least knowledge is a design guideline for developing software, particularly object-oriented programs. In its general form, the LoD is a specific case of loose coupling. The guideline was proposed at Northeastern University towards the end of 1987, and can be succinctly summarized in each of the following ways:

- Each unit should have only limited knowledge about other units: only units "closely" related to the current unit.
- Each unit should only talk to its friends; don't talk to strangers.
- Only talk to your immediate friends.

#### Bad Example

```go
type Money struct {
	Amount int
}

type Wallet struct {
	Cash Money
}

type Person struct {
	Name   string
	Wallet Wallet
}

func (p Person) CanAfford(amount int) bool {
    return p.Wallet.Cash.Amount >= amount
}
```

#### Good Example

```go
type Money struct {
	Amount int
}

func (money Money) CanSpendBy(amount int) bool {
    return money.Amount >= amount
}

type Wallet struct {
	Cash Money
}

func (w Wallet) CanSpendBy(amount int) bool {
    return w.Cash.CanSpendBy(amount)
}

type Person struct {
	Name   string
	Wallet Wallet
}

func (p Person) CanAfford(amount int) bool {
    return p.Wallet.CanSpendBy(amount)
}
```