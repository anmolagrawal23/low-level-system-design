package user

import (
	"errors"
	"fmt"
	"sync/atomic"
)

var	EmptyUserName error = errors.New("EmptyNameError: User name cannot be empty")
var EmptyEmailOrPhone error = errors.New("EmptyEmailOrPhone: Both email and phone number cannot be empty")

var UserID uint32

func AddUser(name, email, phone string) (uint32, error) {
	if len(name) == 0 {
		return 0,EmptyUserName
	}
	if len(email) == 0 && len(phone) == 0 {
		return 0, EmptyEmailOrPhone
	}

	atomic.AddUint32(&UserID,1)
	id := UserID
	newUser := &User{Id: id, Name: name, Email: email, Phone: phone}
	List[id] = newUser
	return id, nil
}

func IfUserPresent(userID uint32) bool {
	_,found := List[userID]
	return found
}

func DeleteUser(userID uint32) error {
	if IfUserPresent(userID) {
		delete(List, userID)
		return nil
	}
	return fmt.Errorf("UserNotFound %d: User not found", userID)
}