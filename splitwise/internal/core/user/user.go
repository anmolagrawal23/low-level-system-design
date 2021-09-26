package user

var List map[uint32]*User
type PendingFriendExpenses map[uint32]float64

type User struct {
	Id		uint32
	Name	string
	Email	string
	Phone	string
}