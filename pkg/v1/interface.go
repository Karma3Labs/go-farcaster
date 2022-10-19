package farcaster_v1

import "github.com/eigentrust-io/go-farcaster/pkg/farcaster"

// API is the Farcaster v1 API.
type API interface {
	// UserRegistrations() returns all user registrations.
	UserRegistrations() UserRegistrationIter
}

// UserRegistrationIter is an iterator over zero or more user registrations.
type UserRegistrationIter interface {
	// Fetch returns (the next user registration record, nil) on success;
	// (nil, nil) if no more records; (nil, error) upon error.
	Fetch() (*farcaster.UserRegistration, error)

	// Stop signals the iterator will no longer be used.
}
