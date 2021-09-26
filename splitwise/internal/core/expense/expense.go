package expense

import (
	"time"
)

type SplitType int
var List []*Expense
var UserExpenseList map[uint32]map[uint32]float64

const (
	Equal SplitType = iota
	Percentage
	Share
)

type Expense struct {
	Id				uint32
	Description		string
	Timestamp		time.Time
	Amount			float64
	SplitType		SplitType
	PaidBy			uint32
	UserShare		map[uint32]float64
}