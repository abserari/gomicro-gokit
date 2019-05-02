package main

import "errors"

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

var ErrEmpty = errors.New("string is empty")
