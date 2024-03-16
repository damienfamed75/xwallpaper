package main

import (
	"errors"
	"strings"
)

type Action uint8

const (
	Set Action = iota
	Apply
	Append
	Get
)

func ActionFromString(s string) (Action, error) {
	switch strings.ToLower(s) {
	case "set":
		return Set, nil
	case "apply":
		return Apply, nil
	case "append":
		return Append, nil
	case "get":
		return Get, nil
	default:
		return 0, errors.New("invalid action")
	}
}
