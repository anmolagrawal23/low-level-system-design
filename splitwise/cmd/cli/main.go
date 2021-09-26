package main

import (
	"bufio"
	"fmt"
	"github.com/anmolagrawal23/low-level-system-design/splitwise/internal/core/expense"
	"github.com/anmolagrawal23/low-level-system-design/splitwise/internal/core/group"
	"github.com/anmolagrawal23/low-level-system-design/splitwise/internal/core/user"
	"os"
	"strconv"
)

func main() {
	Init()
	fmt.Println("Welcome to Splitwise...")

	inputMsg()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter your option: ")
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0 {
			break
		}

		option,err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input:", input)
		}

		switch option {
		case 1:
			fmt.Println("Adding user...")
			AddUser()
		case 2:
			fmt.Println("Adding group...")
			AddGroup()
		case 3:
			fmt.Println("Adding expense...")
			AddExpense()
		case 4:
			fmt.Println("Adding expense in group...")
			AddExpenseByGroup()
		case 5:
			fmt.Println("Printing your balances...")
			PrintMyBalance()
		case 6:
			fmt.Println("Printing all expenses...")
			PrintAllExpenses()
		case 7:
			fmt.Println("Printing group list...")
			PrintGroupList()
		case 8:
			fmt.Println("Printing users list...")
			PrintUserList()
		default:
			inputMsg()
		}
	}
}

func inputMsg() {
	fmt.Println("Enter your input:")
	fmt.Println("1. Add User - Enter name, email and/or phone")
	fmt.Println("2. Add Group - Enter name, comma seperated users' id list")
	fmt.Println("3. Add Expense with users list - Enter amount, split type, payee user id, users share in the form id: amount/percent")
	fmt.Println("4. Add Expense using Group - Enter amount, payee user's id, group id")
	fmt.Println("5. Get your balance list - Enter your user id")
	fmt.Println("6. Print all expenses list")
	fmt.Println("7. Print group list")
	fmt.Println("8. Print Users list")
}

func Init() {
	expense.List = make([]*expense.Expense, 0)
	group.List = make(map[uint32]*group.Group)
	user.List = make(map[uint32]*user.User)
	expense.UserExpenseList = make(map[uint32]map[uint32]float64)
}
