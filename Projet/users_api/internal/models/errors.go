package models

import "fmt"

type CustomError struct {
	Message string `default:""`
	Code    int    `default:"200"`
}

type NotFoundError struct {
	Resource string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%d - %s", e.Code, e.Message)
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.Resource)
}