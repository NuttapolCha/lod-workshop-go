# LoD Mini Workshop Golang

## Introduction

This is a mini workshop for learning the LoD (Law of Demeter) in Golang.
You will learn by doing the exercises in this workshop.

## Setup

1. Clone this repository

   ```sh
   https://github.com/NuttapolCha/lod-workshop-go.git
   ```

2. Checkout to main branch

    ```sh
    git checkout main
    ```

3. You will coding at `./workshop/workshop.go`

## Getting Started

You are given the pre-defined structs and the function signature. You job is to implement the *HandlePersonPurchaseProducts* function that satisfying requirements;

### Requirements

A person can purchase a quantity of products if

1. The person has enough money, i.e. person's money after purchasing product must be greater than or equal to zero.
2. The product is in stock and has enough quantity.

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

Try to implement the *HandlePersonPurchaseProductsWithPromotion* then try to test if your code satisfy the requirements.

```sh
make test
```
