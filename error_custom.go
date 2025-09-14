package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "validation error"}
	}

	if id != "Fitria" {
		return &notFoundError{Message: "data not found"}
	}

	return nil
}

func main() {
	err := SaveData("wonwoo", nil)

	if err != nil {
		/**
		if validationError, ok := err.(*validationError); ok {
			fmt.Println("Validation Error:", validationError.Message)
		} else if notFoundError, ok := err.(*notFoundError); ok {
			fmt.Println("Not Found Error:", notFoundError.Message)
		} else {
			fmt.Println("Other Error:", err.Error())
		}
		*/

		// switch
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation error:", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error:", finalError.Error())
		default:
			fmt.Println("other error:", finalError.Error())
		}
	} else {
		fmt.Println("success")
	}
}
