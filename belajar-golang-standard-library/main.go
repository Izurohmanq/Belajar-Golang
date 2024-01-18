package main

import (
	"fmt"
	"path"
)

// "reflect"

// "time"
// "math"
// "strconv"
// "strings"
// 	"flag"
// 	"os"
// 	"errors"

//package error
// var (
// 	ValidationError = errors.New("Validation Error")
// 	NotFoundError   = errors.New("Not Found Error")
// )

// func GetById(id string) error {
// 	if id == "" {
// 		return ValidationError
// 	}

// 	if id != "asep" {
// 		return NotFoundError
// 	}

// 	return nil
// }

// func main() {
// 	err := GetById("awokawokawo")
// 	if err != nil {
// 		if errors.Is(err, ValidationError) {
// 			fmt.Println("Validation error")
// 		} else if errors.Is(err, NotFoundError) {
// 			fmt.Println("Not Found Errror")
// 		} else {
// 			fmt.Println("Unknown Error")
// 		}
// 	}
// }

// package os
// func main() {
// 	args := os.Args
// 	for _, arg := range args {
// 		fmt.Println(arg)
// 	}

// 	hostname, err := os.Hostname()

// 	if err == nil {
// 		fmt.Println(hostname)
// 	} else {
// 		fmt.Println("Error", err.Error())
// 	}
// }

// package flag
// func main() {
// 	var username *string = flag.String("username", "root", "database username")
// 	var password *string = flag.String("password", "root", "database password")
// 	var host *string = flag.String("host", "root", "database host")
// 	var port *int = flag.Int("port", 0, "database port")

// 	flag.Parse()

// 	fmt.Println("Username", *username)
// 	fmt.Println("password", *password)
// 	fmt.Println("host", *host)
// 	fmt.Println("port", *port)
// }

// package strings
// func main() {
// 	fmt.Println(strings.Contains("Eko Kurniawan", "Eko"))
// 	fmt.Println(strings.Split("Ngentot banget", " "))
// 	fmt.Println(strings.ToUpper("Ngentot banget"))
// 	fmt.Println(strings.ToLower("Ngentot banget"))
// 	fmt.Println(strings.Trim("Ngentot           banget        ", " "))
// 	fmt.Println(strings.ReplaceAll("Ngentot banget", "Ngentot", "Bagus"))
// }

// package strconv
// func main() {
// 	boolean, err := strconv.ParseBool("true")

// 	if err != nil {
// 		fmt.Println("error", err.Error())
// 	} else {
// 		fmt.Println(boolean)
// 	}

// 	resultInt, err := strconv.Atoi("1000")

// 	if err != nil {
// 		fmt.Println("error", err.Error())
// 	} else {
// 		fmt.Println(resultInt)
// 	}

// 	binary := strconv.FormatInt(999, 2)
// 	fmt.Println(binary)

// 	stringInt := strconv.Itoa(9999)
// 	fmt.Println(stringInt)
// }

// package math
// func main() {
// 	fmt.Println(math.Ceil(1.40))
// 	fmt.Println(math.Floor(1.40))
// 	fmt.Println(math.Round(1.50))
// 	fmt.Println(math.Max(1.40, 100))
// 	fmt.Println(math.Min(1.40, 100))
// }

//package list/container buat linked list
// func main() {
// 	var data *list.List = list.New()

// 	data.PushBack("Asep")
// 	data.PushBack("Kusen")
// 	data.PushBack("Kayu")

// 	var head *list.Element = data.Front()
// 	fmt.Println(head.Value)

// 	next := head.Next()
// 	fmt.Println(next.Value)

// 	next = next.Next()
// 	fmt.Println(next.Value)

// 	for i := data.Front(); i != nil; i = i.Next() {
// 		fmt.Println(i.Value)
// 	}
// }

//package sort
// type User struct {
// 	Name string
// 	Age  int
// }

// type UserSlice []User

// func (s UserSlice) Len() int {
// 	return len(s)
// }

// func (s UserSlice) Less(i, j int) bool {
// 	return s[i].Age < s[j].Age
// }

// func (s UserSlice) Swap(i, j int) {
// 	s[i], s[j] = s[j], s[i]
// }

// func main() {
// 	users := []User{
// 		{"Eko", 30},
// 		{"Budi", 20},
// 		{"Asep", 9},
// 		{"Buadu", 99},
// 	}

// 	sort.Sort(UserSlice(users))

// 	fmt.Println(users)
// }

//time
// func main() {
// 	var now time.Time = time.Now()
// 	fmt.Println(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())

// 	var utc time.Time = time.Date(2009, time.May, 17, 0, 0, 0, 0, time.UTC)

// 	fmt.Println(utc)
// }

// reflect
// type Sample struct {
// 	Name string `required:"true" max:"10"`
// }

// type Person struct {
// 	Name string `required:"true" max:"10"`
// }

// func readField(value any) {
// 	var valueType reflect.Type = reflect.TypeOf(value)

// 	fmt.Println("Type name", valueType.Name())

// 	for i := 0; i < valueType.NumField(); i++ {
// 		var structField reflect.StructField = valueType.Field(i)
// 		fmt.Println(structField.Name, "with type", structField.Type)
// 		fmt.Println(structField.Tag.Get("required"))
// 		fmt.Println(structField.Tag.Get("max"))
// 	}
// }

// func IsValid(value any) (result bool) {
// 	t := reflect.TypeOf(value)
// 	for i := 0; i < t.NumField(); i++ {
// 		f := t.Field(i)
// 		if f.Tag.Get("required") == "true" {
// 			data := reflect.ValueOf(value).Field(i).Interface()
// 			result = data != ""
// 			if result == false {
// 				return result
// 			}
// 		}
// 	}
// 	return result
// }

// func main() {
// 	readField(Sample{"koko"})
// 	readField(Person{"koko"})

// 	person := Person{
// 		Name: "padli",
// 	}
// 	fmt.Println(IsValid(person))
// }

//regexp
// func main() {
// 	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

// 	fmt.Println(regex.MatchString("koko"))
// 	fmt.Println(regex.MatchString("kpk"))
// 	fmt.Println(regex.MatchString("Konco"))
// 	fmt.Println(regex.MatchString("eko"))
// }

// encoding
// func main() {
// 	value := "Fadli Izurohman"

// 	encoded := base64.StdEncoding.EncodeToString([]byte(value))

// 	fmt.Println(encoded)

// 	decoded, err := base64.StdEncoding.DecodeString(encoded)

// 	if err != nil {
// 		fmt.Println("Error", err.Error())
// 	} else {
// 		fmt.Println(string(decoded))
// 	}

// 	cvstring := "Fadli,Izurohman,Kasep\n" + "Windah,Batubata,oakwoakwoa\n" + "Raisya,adrianada,kue\n" + "Caca,marica,hehey\n"

// 	reader := csv.NewReader(strings.NewReader(cvstring))

// 	for {
// 		record, err := reader.Read()

// 		if err == io.EOF {
// 			break

// 		}

// 		fmt.Println(record)
// 	}

// 	writer := csv.NewWriter(os.Stdout)

// 	_ = writer.Write([]string{"aowkawok", "oakwdoawok", "asep"})
// 	_ = writer.Write([]string{"iya", "aasdadw", "papa"})
// 	_ = writer.Write([]string{"masa", "iyadeh", "ywdh"})

// 	writer.Flush()
// }

// slices
// func main() {
// 	names := []string{"john", "asep", "kusen"}
// 	values := []int{100, 23, 323}

// 	fmt.Println(slices.Min(names))
// 	fmt.Println(slices.Min(values))
// 	fmt.Println(slices.Max(names))
// 	fmt.Println(slices.Max(values))
// }

//path

func main() {
	fmt.Println(path.Dir("aowkaokok/world.go"))
	fmt.Println(path.Base("aowkaokok/world.go"))
	fmt.Println(path.Ext("aowkaokok/world.go"))
	fmt.Println(path.Join("hello", "world", "main.go"))
}
