package ozzo

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&u.Email, validation.Required, validation.Email),
		validation.Field(&u.Age, validation.Required, validation.Min(18)), // tip is here
	)
}

func main() {
	user := User{
		Name:  "Ali",
		Email: "ali@example.com",
		Age:   17,
	}

	if err := user.Validate(); err != nil {
		fmt.Println("Validation failed:", err)
	} else {
		fmt.Println("Validation succeeded!")
	}
}

// ali is under 18 -----> failing validation
