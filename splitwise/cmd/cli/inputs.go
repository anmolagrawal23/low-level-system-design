package main

import (
	"bufio"
	"fmt"
	"github.com/anmolagrawal23/low-level-system-design/splitwise/internal/core/expense"
	"github.com/anmolagrawal23/low-level-system-design/splitwise/internal/core/group"
	"github.com/anmolagrawal23/low-level-system-design/splitwise/internal/core/user"
	"os"
	"strconv"
	"strings"
)

func AddUser() {
	scanner := bufio.NewScanner(os.Stdin)
	var userName, email, phone string

	fmt.Print("Enter User Name: ")
	scanner.Scan()
	userName = scanner.Text()

	fmt.Print("Enter email id: ")
	scanner.Scan()
	email = scanner.Text()

	fmt.Print("Enter Phone number: ")
	scanner.Scan()
	phone = scanner.Text()

	uID, err := user.AddUser(userName, email, phone)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User created with ID:", uID)
	}
}

func AddGroup() {
	scanner := bufio.NewScanner(os.Stdin)
	var groupName, users string

	fmt.Print("Enter Group Name: ")
	scanner.Scan()
	groupName = scanner.Text()

	fmt.Print("Enter comma separated user ids: ")
	scanner.Scan()
	users = scanner.Text()

	userStrList := strings.Split(users, ",")
	userList := make([]uint32, 0)

	for _,user := range userStrList {
		userId, err := strconv.ParseUint(user, 10, 32)
		if err != nil {
			panic(fmt.Sprint("Invalid user id:", user))
		} else {
			userList = append(userList, uint32(userId))
		}
	}
	gID, err := group.CreateGroup(groupName, userList)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Group created with ID:", gID)
	}
}

func AddExpense() {
	scanner := bufio.NewScanner(os.Stdin)
	var amount float64
	var description string
	var splitType expense.SplitType
	var paidByUId uint32
	var share string

	shareRatio := make(map[uint32]float64)
	var input string
	var err error

	fmt.Print("Enter expense amount: ")
	scanner.Scan()
	input = scanner.Text()
	amount, err = strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Error parsing amount...")
	}

	fmt.Print("Enter user who paid: ")
	scanner.Scan()
	input = scanner.Text()
	var tmpUint64 uint64
	tmpUint64, err = strconv.ParseUint(input, 10, 32)
	paidByUId = uint32(tmpUint64)

	fmt.Print("Enter split type (0-Equal, 1-Percentage, 2-Share amount): ")
	scanner.Scan()
	input = scanner.Text()
	var inputSplitType int
	inputSplitType, err = strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error parsing split type...")
	}
	splitType = expense.SplitType(inputSplitType)

	if splitType == expense.Equal {
		fmt.Print("Enter comma seperated list of user id: ")
		scanner.Scan()
		share = scanner.Text()

		userList := strings.Split(strings.TrimSpace(share), ",")
		for _,user := range userList {
			userId, err := strconv.ParseUint(user, 10, 32)
			if err != nil {
				panic(fmt.Sprint("Invalid user id:", user))
			}
			shareRatio[uint32(userId)] = 0
		}
	} else {

		fmt.Print("Enter list of users in the form of 'userID1:share1, userID2:share2,...': ")
		scanner.Scan()
		share = scanner.Text()
		shareList := strings.Split(share, ",")
		for _,userShareStr := range shareList {
			userShare := strings.Split(userShareStr, ":")
			userId, err := strconv.ParseUint(strings.TrimSpace(userShare[0]), 10, 32)
			if err != nil {
				panic(fmt.Sprint("Invalid user id:", userShare[0]))
			}
			amount, err := strconv.ParseFloat(strings.TrimSpace(userShare[1]), 64)
			if err != nil {
				panic(fmt.Sprint("Invalid share/percent:", userShare[1]))
			}
			shareRatio[uint32(userId)] = amount
		}
	}

	fmt.Print("Enter description for the expense: ")
	scanner.Scan()
	description = scanner.Text()

	err = expense.CreateNewExpense(amount, description, splitType, paidByUId, shareRatio)
	if err != nil {
		fmt.Println(err)
	}
}

func AddExpenseByGroup()  {
	var amount float64
	var description string
	var groupID uint32
	var paidByUId uint32
	var input string
	var err error
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter expense amount: ")
	scanner.Scan()
	input = scanner.Text()
	amount, err = strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Error parsing amount...")
	}

	fmt.Print("Enter user who paid: ")
	scanner.Scan()
	input = scanner.Text()
	var tmpUint64 uint64
	tmpUint64, err = strconv.ParseUint(input, 10, 32)
	paidByUId = uint32(tmpUint64)

	fmt.Print("Enter the group id: ")
	scanner.Scan()
	input = scanner.Text()
	tmpUint64, err = strconv.ParseUint(input, 10, 32)
	if err != nil {
		fmt.Println("Error parsing group id...")
	}
	groupID = uint32(tmpUint64)

	fmt.Print("Enter description for the expense: ")
	scanner.Scan()
	description = scanner.Text()

	err = expense.CreateExpenseByGroup(amount, description, paidByUId, groupID)
	if err != nil {
		fmt.Println(err)
	}
}

func PrintMyBalance() {
	var uID uint32
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter user who paid: ")
	scanner.Scan()
	input := scanner.Text()
	tmpUint64, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		fmt.Println("Error parsing user id...")
	}
	uID = uint32(tmpUint64)
	if userList, found := expense.UserExpenseList[uID]; !found {
		fmt.Println("You are all settled up!")
	} else {
		for user,amount := range userList {
			fmt.Println(user, amount)
		}
	}
}

func PrintAllExpenses() {
	fmt.Printf("ExpenseID\tAmount\tPaidBy\tUserShare\tTimestamp\n")
	for _,exp := range expense.List {
		fmt.Printf("%d\t%f\t%d\t%v\t%v\n", exp.Id, exp.Amount, exp.PaidBy, exp.UserShare, exp.Timestamp)
	}
}

func PrintUserList() {
	fmt.Printf("UserID\tName\tPhone\tEmail\n")
	for uID, user := range user.List {
		fmt.Printf("%d\t%s\t%s\t%s\n", uID, user.Name, user.Phone, user.Email)
	}
}

func PrintGroupList() {
	fmt.Printf("GroupID\tName\tUsers List\n")
	for gID, gInfo := range group.List {
		fmt.Printf("%d\t%s\t%v\n",gID, gInfo.Name, gInfo.UserList)
	}
}