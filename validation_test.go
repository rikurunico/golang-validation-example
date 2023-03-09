package golangvalidation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidatePath(t *testing.T) {
	validate := validator.New()

	if validate == nil {
		t.Error("validate is nil")
	}
}

func TestValidationField(t *testing.T) {
	validate := validator.New()
	var user string = "Nico"

	err := validate.Var(user, "required")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationTwoVariable(t *testing.T) {
	validate := validator.New()

	password := "123456"
	password2 := "123456"

	err := validate.VarWithValue(password, password2, "eqfield")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMultipleTag(t *testing.T) {
	validate := validator.New()
	var email string = "nico@gmail.com"

	err := validate.Var(email, "required,email")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestTagParameter(t *testing.T) {
	validate := validator.New()
	user := "12345"

	err := validate.Var(user, "required,numeric,min=5,max=10")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestStruct(t *testing.T) {
	validate := validator.New()

	type User struct {
		Email    string `validate:"required,email"`
		Password string `validate:"required,min=6,max=10"`
	}

	user := User{
		Email:    "nico@wazirnico.com",
		Password: "123456",
	}

	err := validate.Struct(user)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationError := range validationErrors {
			fmt.Println("error", validationError.Field(), "on tag", validationError.Tag(), "with error", validationError.Error())
		}
	}
}
