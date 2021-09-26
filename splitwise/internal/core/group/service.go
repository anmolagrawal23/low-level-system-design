package group

import (
	"errors"
	"fmt"
	user2 "github.com/anmolagrawal23/low-level-system-design/splitwise/internal/core/user"
	"sync/atomic"
	"time"
)

var GroupID uint32 = 0
var	EmptyGroupName error = errors.New("EmptyNameError: Group name cannot be empty")

func CreateGroup(name string, userList []uint32) (uint32,error) {

	if len(name) == 0 {
		return 0, EmptyGroupName
	}

	atomic.AddUint32(&GroupID, 1)
	id := GroupID
	userMap := make(map[uint32]bool)

	for _,user := range userList {
		if !user2.IfUserPresent(user) {
			return 0, fmt.Errorf("UserNotFound %d: user not found", user)
		}
		userMap[user] = true
	}

	newGroup := &Group{Id: id, Name: name, UserList: userMap, CreatedTime: time.Now()}
	List[id] = newGroup

	return id, nil
}
