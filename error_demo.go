package main 

import "fmt"

type MyError struct {
	message string 
}

func (err *MyError) Error() string {
	return fmt.Sprintf("Error is %s", err.message)
}

func run() error {
	return &MyError{"Running is not supported"}
}

func main() {
	if err:=run(); err != nil {
		fmt.Println(err)
	}
}