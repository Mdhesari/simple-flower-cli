package main

import (
	"bufio"
	"fmt"
	"mdhesari/coralflora/model"
	"mdhesari/coralflora/repositories"
	"os"
	"strconv"
)

var (
	scanner    *bufio.Scanner = bufio.NewScanner(os.Stdin)
	id         int
	name       string
	price      int
	stockCount int
	flower     *model.Flower
	flowers    []*model.Flower
)

func main() {
	fmt.Println("Welcome To Coralflora!")

	flowers = repositories.GetItems()

	for {
		fmt.Println("Please enter a command : ")
		scanner.Scan()
		command := scanner.Text()

		switch command {
		case "buy":
			// buy flower
			buyFlower()
		case "add":
			// add flower
			if !isAdmin() {
				fmt.Println("Password is invalid!")

				continue
			}

			addFlower()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid command.")
		}
	}
}

func isAdmin() bool {
	password := askQuestion("Please enter admin password :")

	return password == "secret"
}

func addFlower() {
	name = askQuestion("Enter the name :")

	isFound, _ := repositories.FindByName(flowers, name)

	if (isFound != nil) {
		fmt.Println("Name already exists!")
	}

	price, _ = strconv.Atoi(askQuestion("Enter the price :"))
	stockCount, _ = strconv.Atoi(askQuestion("Enter stock count :"))

	f := model.Flower{
		Name:       name,
		Price:      price,
		StockCount: stockCount,
	}

	insertIntoDataStorage(&f)
}

func insertIntoDataStorage(flower *model.Flower) {
	flower.Id = len(flowers) + 1
	flowers = append(flowers, flower)

	repositories.UpdateItems(flowers)
}

func askQuestion(q string) string {
	fmt.Println(q)
	scanner.Scan()

	return scanner.Text()
}

func buyFlower() {
	id, _ = strconv.Atoi(askQuestion("Enter the flower Id :"))

	flower, id = repositories.FindById(flowers, id)

	if flower == nil {
		fmt.Println("Flower is not found!")

		return
	}

	if flower.StockCount < 1 {
		fmt.Println("No flower in stock.")

		return
	}

	flower.StockCount -= 1
	flowers[id] = flower

	repositories.UpdateItems(flowers)

	fmt.Println("Thanks for buying our flower!")
	fmt.Printf("Now there is only %d in stock.", flower.StockCount)
	fmt.Println()
}