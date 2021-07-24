package entity

import "errors"

//ErrNotFound not found
var NoError error = nil

//ErrNotFound not found
var ErrNotFound = errors.New("not found")

//ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("invalid entity")

//ErrCannotBeDeleted cannot be deleted
var ErrCannotBeDeleted = errors.New("cannot be deleted")

//ErrApplAlreadyExist application already exists in the database
var ErrAlreadyExist = errors.New("already exists")

//ErrInvalidEntity invalid entity
var ErrInvalidCredentials = errors.New("invalid credentials")

//ErrInvalidTokenEntity invalid entity
var ErrInvalidToken = errors.New("invalid token")

//ErrInvalidTokenEntity invalid entity
var ErrWarningToken = errors.New("warning! almost out of token")

// //ErrNotEnoughBooks cannot borrow
// var ErrNotEnoughBooks = errors.New("Not enough books")

// //ErrBookAlreadyBorrowed cannot borrow
// var ErrBookAlreadyBorrowed = errors.New("Book already borrowed")

// //ErrBookNotBorrowed cannot return
// var ErrBookNotBorrowed = errors.New("Book not borrowed")

var ErrCodeMapper = map[error]int{
	ErrNotFound:           10,
	ErrInvalidEntity:      20,
	ErrCannotBeDeleted:    30,
	ErrAlreadyExist:       40,
	ErrInvalidCredentials: 50,
	ErrInvalidToken:       51,
}
