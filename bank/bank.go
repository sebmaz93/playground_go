package bank

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFileName = "balance.txt"

func readFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, errors.New("failed to read from file")
	}
	balanceTxt := string(data)
	balance, err := strconv.ParseFloat(balanceTxt, 64)
	if err != nil {
		return 0, errors.New("failed to parse")
	}
	return balance, nil
}

func writeToFile(fileName string, balance float64) {
	balanceTxt := fmt.Sprint(balance)
	err := os.WriteFile(fileName, []byte(balanceTxt), 0644)
	if err != nil {
		return
	}
}

func main() {
	var balance, err = readFromFile(balanceFileName)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------")
		panic(err)
	}

	for {
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1 check balance")
		fmt.Println("2 deposit")
		fmt.Println("3 withdraw")
		fmt.Println("4 exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("your balance is", balance)
		} else if choice == 2 {
			var deposit float64
			fmt.Print("enter deposit amount ")
			fmt.Scan(&deposit)
			balance += deposit
			writeToFile(balanceFileName, balance)
			fmt.Println("new total amount: ", balance)
		} else if choice == 3 {
			var withdraw float64
			fmt.Print("enter withdraw amount ")
			fmt.Scan(&withdraw)
			balance -= withdraw
			writeToFile(balanceFileName, balance)
			fmt.Println("new total amount: ", balance)
		} else {
			fmt.Println("Bye")
			//return
			break
		}

		fmt.Println("choice: ", choice)

	}

}
