package utils

import "fmt"

type ValidateError struct {
	Key string
	Msg string
}

type Errors struct {
	List []ValidateError
}

func NewErrors() *Errors {
	return &Errors{}
}

func (e *Errors) Add(key, value string) {
	e.List = append(e.List, ValidateError{Key: key, Msg: value})
}

func (e *Errors) HasError() bool {
	return len(e.List) > 0
}

func (e *Errors) GetErrorsString() string {
	errorStr := ""
	for _, e := range e.List {
		errorStr += fmt.Sprintf("%s: %s\n", e.Key, e.Msg)
	}
	return errorStr
}
