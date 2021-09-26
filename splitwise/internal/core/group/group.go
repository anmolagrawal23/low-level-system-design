package group

import (
	"time"
)

var List map[uint32]*Group

type Group struct {
	Id			uint32
	Name		string
	UserList	map[uint32]bool
	CreatedTime	time.Time
}