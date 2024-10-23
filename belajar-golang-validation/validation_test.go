package belajargolangvalidation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

// pake validator

func TestValidation(t *testing.T) {
	var validate *validator.Validate = validator.New()
	if validate == nil {
		t.Error("Validate is Nill")
	}
}

func TestValidationVariable(t *testing.T) {
	validate := validator.New()
	user := "Asep"

	err := validate.Var(user, "required")

	if err != nil {
		fmt.Println(err.Error())
	}

}

func TestValidationTwoVariables(t *testing.T) {
	validate := validator.New()

	password := "rahasia"
	confirmPassword := "rahasia"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")

	if err != nil {
		fmt.Println(err.Error())
	}

}

func TestMultipleTag(t *testing.T) {
	validate := validator.New()
	user := "asep"

	err := validate.Var(user, "required,numeric") // setelah sesudah koma gak boleh spasi

	if err != nil {
		fmt.Println(err.Error())
	}

}

func TestTagParameter(t *testing.T) {
	validate := validator.New()
	user := "99123"

	err := validate.Var(user, "required,numeric,min=5,max=10") // setelah sesudah koma gak boleh spasi

	if err != nil {
		fmt.Println(err.Error())
	}

}

func TestStruct(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"` // tetap gakboleh ada spasi
		Password string `validate:"required,min=5"`
	}
	validate := validator.New()

	loginRequest := LoginRequest{
		Username: "Fadli@gamil.com",
		Password: "Fadli123",
	}

	err := validate.Struct(loginRequest)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationErrors(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"` // tetap gakboleh ada spasi
		Password string `validate:"required,min=5"`
	}
	validate := validator.New()

	loginRequest := LoginRequest{
		Username: "Fadli",
		Password: "Fadli123",
	}

	err := validate.Struct(loginRequest)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

func TestStructCrossField(t *testing.T) {
	type RegisterUser struct {
		Username        string `validate:"required,email"` // tetap gakboleh ada spasi
		Password        string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
	}
	validate := validator.New()

	registerUser := RegisterUser{
		Username:        "Fadli@gamil.com",
		Password:        "Fadli123",
		ConfirmPassword: "Fadli123",
	}

	err := validate.Struct(registerUser)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestNestedStruct(t *testing.T) {
	type Address struct {
		City    string `validate:"required"` // tetap gakboleh ada spasi
		Country string `validate:"required"`
	}
	type User struct {
		Id        string  `validate:"required"` // tetap gakboleh ada spasi
		Name      string  `validate:"required"`
		Addresses Address `validate:"required"`
	}

	validate := validator.New()

	registerUser := User{
		Id:   "UBUBUBUB",
		Name: "Fadli123",
		Addresses: Address{
			City:    "",
			Country: "",
		},
	}

	err := validate.Struct(registerUser)

	if err != nil {
		fmt.Println(err.Error())
	}
}
func TestCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"` // tetap gakboleh ada spasi
		Country string `validate:"required"`
	}
	type User struct {
		Id        string    `validate:"required"` // tetap gakboleh ada spasi
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
	}

	validate := validator.New()

	registerUser := User{
		Id:   "UBUBUBUB",
		Name: "Fadli123",
		Addresses: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
	}

	err := validate.Struct(registerUser)

	if err != nil {
		fmt.Println(err.Error())
	}
}
func TestBasicCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"` // tetap gakboleh ada spasi
		Country string `validate:"required"`
	}
	type User struct {
		Id        string    `validate:"required"` // tetap gakboleh ada spasi
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
		Hobbies   []string  `validate:"dive,required,min=3"`
	}

	validate := validator.New()

	registerUser := User{
		Id:   "UBUBUBUB",
		Name: "Fadli123",
		Addresses: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
		Hobbies: []string{
			"Gaming",
			"Malas",
			"ee",
		},
	}

	err := validate.Struct(registerUser)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMap(t *testing.T) {
	type Address struct {
		City    string `validate:"required"` // Tidak boleh kosong
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"` // Tidak boleh kosong
	}

	type User struct {
		Id        string            `validate:"required"` // Tidak boleh kosong
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"dive,required,min=3"`
		School    map[string]School `validate:"required,dive,keys,required,min=2,endkeys,required"` // Validasi kunci dan nilai
	}

	validate := validator.New()

	registerUser := User{
		Id:   "UBUBUBUB",
		Name: "Fadli123",
		Addresses: []Address{
			{
				City:    "Jakarta", // Harus ada nilai untuk validasi berhasil
				Country: "Indonesia",
			},
			{
				City:    "Bandung",
				Country: "Indonesia",
			},
		},
		Hobbies: []string{
			"Gaming", // Lebih dari 3 karakter
			"Malas",  // Lebih dari 3 karakter
			"BAB",    // Tepat 3 karakter, valid
		},
		School: map[string]School{
			"SD": {
				Name: "Al Amanah", // Valid
			},
			"SMP": {
				Name: "Ciguruik", // Valid
			},
			"": {
				Name: "",
			},
		},
	}

	err := validate.Struct(registerUser)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Validation success!")
	}
}
func TestBasicMap(t *testing.T) {
	type Address struct {
		City    string `validate:"required"` // Tidak boleh kosong
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"` // Tidak boleh kosong
	}

	type User struct {
		Id        string            `validate:"required"` // Tidak boleh kosong
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"dive,required,min=3"`
		School    map[string]School `validate:"required,dive,keys,required,min=2,endkeys,required"` // Validasi kunci dan nilai
		Wallet    map[string]int    `validate:"required,dive,keys,required,endkeys,required,gt=0"`  // Validasi kunci dan nilai
	}

	validate := validator.New()

	registerUser := User{
		Id:   "UBUBUBUB",
		Name: "Fadli123",
		Addresses: []Address{
			{
				City:    "Jakarta", // Harus ada nilai untuk validasi berhasil
				Country: "Indonesia",
			},
			{
				City:    "Bandung",
				Country: "Indonesia",
			},
		},
		Hobbies: []string{
			"Gaming", // Lebih dari 3 karakter
			"Malas",  // Lebih dari 3 karakter
			"BAB",    // Tepat 3 karakter, valid
		},
		School: map[string]School{
			"SD": {
				Name: "Al Amanah", // Valid
			},
			"SMP": {
				Name: "Ciguruik", // Valid
			},
			"": {
				Name: "",
			},
		},
		Wallet: map[string]int{
			"BCA":     0,
			"Mandiri": 1000000,
			"":        9,
		},
	}

	err := validate.Struct(registerUser)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Validation success!")
	}
}

func TestAlias(t *testing.T) {
	validate := validator.New()
	validate.RegisterAlias("varchar", "required,max=255")

	type Seller struct {
		Id     string `validate:"varchar"`
		Name   string `validate:"varchar"`
		Owner  string `validate:"varchar"`
		Slogan string `validate:"varchar"`
	}

	seller := Seller{
		Id:     "",
		Name:   "",
		Owner:  "",
		Slogan: "",
	}

	err := validate.Struct(seller)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func MustValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value != strings.ToUpper(value) {
			return false
		}

		if len(value) < 5 {
			return false
		}
	}

	return true
}

func TestCustomeValidationFunction(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", MustValidUsername)

	type LoginRequest struct {
		Username string `validate:"required,username"`
		Password string `validate:"required"`
	}

	request := LoginRequest{
		Username: "KOKONTOLAN",
		Password: "",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

var regexNumber = regexp.MustCompile("^[0-9]+$")

func MustValidPin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())

	if err != nil {
		panic(err)
	}

	value := field.Field().String()
	if regexNumber.MatchString(value) {
		return false
	}

	return len(value) == length

}

func TestCustomValidationParameter(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("pin", MustValidPin)

	type Login struct {
		Phone string `validate:"required,number"`
		Pin   string `validate:"required,pin=6"`
	}

	request := Login{
		Phone: "08972330",
		Pin:   "12332",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestOrRule(t *testing.T) {
	type Login struct {
		Username string `validate:"required,email|numeric"`
		Password string `validate:"required"`
	}

	request := Login{
		Username: "213123",
		Password: "sbubu",
	}

	validate := validator.New()
	err := validate.Struct(request)

	if err != nil {
		fmt.Println(err.Error())
	}

}

func MustEqualsIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()
	if !ok {
		panic("Field not oK")
	}

	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestCrossFieldValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("field_equals_ignore_case", MustEqualsIgnoreCase)

	type User struct {
		Username string `validate:"required,field_equals_ignore_case=Email|field_equals_ignore_case=Phone"`
		Email    string `validate:"required,email"`
		Phone    string `validate:"required,numeric"`
		Name     string `validate:"required"`
	}

	user := User{
		Username: "089667899340",
		Email:    "asep@mail.con",
		Phone:    "089667899340",
		Name:     "asep",
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

type RegisterRequest struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required,numeric"`
	Password string `validate:"required"`
}

func MustValidRegisterSuccess(level validator.StructLevel) {
	registerRequest := level.Current().Interface().(RegisterRequest)

	if registerRequest.Username == registerRequest.Email || registerRequest.Username == registerRequest.Phone {

	} else {
		level.ReportError(registerRequest.Username, "Username", "Username", "Username", "")
	}
}

func TestStructLevelValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(MustValidRegisterSuccess, RegisterRequest{})

	request := RegisterRequest{
		Username: "089667899340",
		Email:    "asep@mail.con",
		Phone:    "089667899340",
		Password: "asep",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}
