package expense

import (
	"errors"
	"fmt"
	"github.com/anmolagrawal23/low-level-system-design/splitwise/internal/core/group"
	user2 "github.com/anmolagrawal23/low-level-system-design/splitwise/internal/core/user"
	"sync/atomic"
	"time"
)

var ExpenseID uint32
var InvalidShareRatio error = errors.New("InvalidShareRatio: Share amount or percent is not valid")

func CreateNewExpense(amount float64, description string, splitType SplitType, paidBy uint32, shareRatio map[uint32]float64) error {

	fmt.Printf("Amount: %.2f, PaidBy: %d, Share: %v, Split type: %d, Description: %s\n", amount, paidBy, shareRatio, splitType, description)

	for userID := range shareRatio {
		if !user2.IfUserPresent(userID) {
			return fmt.Errorf("UserNotFound %d: user not found", userID)
		}
	}

	if splitType != Equal {
		if !verifyShareRatio(splitType, amount, shareRatio) {
			return InvalidShareRatio
		}
	}

	if splitType == Percentage {
		shareRatio = calculateShare(amount, shareRatio)
	} else if splitType == Equal {
		shareRatio = calculateEqualShare(amount, shareRatio)
	}

	atomic.AddUint32(&ExpenseID, 1)
	id := ExpenseID
	newExpense := &Expense{	Id: id,
							PaidBy: paidBy,
							Amount: amount,
							UserShare: shareRatio,
							Description: description,
							Timestamp: time.Now(),
							SplitType: splitType}

	List = append(List, newExpense)
	newExpense.updateUserExpenses()

	return nil
}

func CreateExpenseByGroup(amount float64, description string, paidBy uint32, groupID uint32) error {

	if _, found := group.List[groupID]; !found {
		return fmt.Errorf("GroupNotFound %d: group not found", groupID)
	}
	shareRatio := make(map[uint32]float64)
	for uId,_ := range group.List[groupID].UserList {
		shareRatio[uId] = 0.0
	}

	return CreateNewExpense(amount, description, Equal, paidBy, shareRatio)
}

func (ex *Expense) updateUserExpenses()  {

	_, found := UserExpenseList[ex.PaidBy]
	if !found {
		UserExpenseList[ex.PaidBy] = make(map[uint32]float64)
	}

	for uId, amt := range ex.UserShare {
		updateFriendShare(amt, ex.PaidBy, uId)
		updateFriendShare(-1 * amt, uId, ex.PaidBy)
	}
}


func verifyShareRatio(splitType SplitType, amount float64, shareRatio map[uint32]float64) bool {
	if splitType == Share {
		var shareAmt float64
		for _,amt := range shareRatio {
			shareAmt += amt
		}
		if shareAmt != amount {
			return false
		}
	} else {
		var percent float64
		for _,per := range shareRatio {
			percent += per
		}
		if percent != float64(100) {
			return false
		}
	}
	return true
}

func calculateShare(amount float64, shareRatio map[uint32]float64) map[uint32]float64 {
	var lastUser uint32
	var amountSoFar float64

	for uId, percent := range shareRatio {
		shareAmt := (percent/100) * amount
		shareRatio[uId] = shareAmt
		amountSoFar += shareAmt
		lastUser = uId
	}

	if amount-amountSoFar != 0.0 {
		shareRatio[lastUser] += amount-amountSoFar
	}
	return shareRatio
}

func calculateEqualShare(amount float64, shareRatio map[uint32]float64) map[uint32]float64 {

	eqAmt := amount/float64(len(shareRatio))
	for uID,_ := range shareRatio {
		shareRatio[uID] = eqAmt
	}
	return shareRatio
}

func updateFriendShare(amount float64, paidTo, paidBy uint32) {

	if _, found := UserExpenseList[paidBy]; !found {
		UserExpenseList[paidBy] = make(map[uint32]float64)
	}
	_, found := UserExpenseList[paidBy][paidTo]
	if found {
		UserExpenseList[paidBy][paidTo] -= amount
	} else {
		UserExpenseList[paidBy][paidTo] = -1 * amount
	}
}