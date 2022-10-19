package farcaster

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// Username is a Farcaster username.
type Username [32]byte

// NewUsername returns a username from the given string.
func NewUsername(name string) (username Username, err error) {
	if len(name) <= len(username) {
		err = errors.New("username too long")
	} else {
		copy(username[:], name)
	}
	return
}

// MustNewUsername does the same as NewUsername, except it panics upon error.
func MustNewUsername(name string) Username {
	username, err := NewUsername(name)
	if err != nil {
		panic(err)
	}
}

// String returns the string representation of the given username.
func (un *Username) String() string {
	return strings.TrimRight(string(un), "\000")
}

// UserRegistration is a Farcaster user registration record.
type UserRegistration struct {
	Username Username
	Owner    common.Address
}
