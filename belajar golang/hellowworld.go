package main

import (
	"fmt"
)

//	func main() {
//		for i := 0; i < 10; i++ {
//			fmt.Println("aokwoakoakwk")
//		}
//	}

// Number
// func main() {
// 	fmt.Println("Satu = ", 1)
// }

// variable
// func main() {
// 	// := biar lebih sederhana gak usah mendeklarasikan string, dalam arti harus variable tersebut harus baru
// 	kata := "aowkaow awokaow aokawok"

// 	fmt.Println(kata)
// 	fmt.Println(len(kata))
// 	fmt.Println(kata[1])

// 	// cara yang lain
// 	var (
// 		firstname = "kusen"
// 		lastname  = "kayu"
// 	)

// 	fmt.Println(firstname)
// 	fmt.Println(lastname)

// }

// variable const aka constant, inget gak boleh diganti
// func main() {
// 	const firstname = "asep"
// 	const lastname = "kayu"

// 	const (
// 		namaDepan    = "ujang"
// 		namaBelakang = "kusen"
// 	)

// 	fmt.Println(firstname, lastname)
// 	fmt.Println(namaDepan, namaBelakang)
// }

//konversi nilai
// func main() {
// 	var nilai32 int32 = 32768
// 	var nilai64 int64 = int64(nilai32)
// 	var nilai16 int16 = int16(nilai32)

// 	fmt.Println(nilai32)
// 	fmt.Println(nilai64)
// 	fmt.Println(nilai16)

// 	nama := "asep kusen"

//		fmt.Println(string(nama[1]))
//	}

// operasi matematika
// func main() {
// 	angka1 := 1
// 	angka2 := 2

// 	jumlah := angka1 + angka2

// 	fmt.Println("Angka2 = ", jumlah)

// 	angka3 := 3
// 	angka3 += 10

// 	fmt.Println("Angka3 = ", angka3)

// }

// perbandingan
// func main() {
// 	nilai := 91

// 	if nilai%2 == 0 {
// 		fmt.Println("Nilai kamu Genap")
// 	} else {
// 		fmt.Println("Nilai Kamu ganjil")
// 	}
// }

// slice
// func main() {
// 	name := [...]string{"asep", "kusen", "kayu", "oakwoakwoakw"} //ieu array yeuh

// 	slice := name[1:3]
// 	fmt.Println(name)
// 	fmt.Println(slice)
// 	fmt.Println(slice[1])

// 	name2 := append(slice, "oakwokaw")

// 	fmt.Println(name2)

// 	fromSlice := name[:]
// 	toSlice := make([]string, len(fromSlice), cap(fromSlice))

// 	copy(toSlice, fromSlice)

// 	fmt.Println(fromSlice)
// 	fmt.Println(toSlice)
// }

// map
// func main() {
// 	person := map[string]string{
// 		"nama":   "Asep",
// 		"alamat": "aokwoakwoakw",
// 	}

// 	fmt.Println(person)
// 	fmt.Println(person["nama"])
// 	fmt.Println(person["alamat"])

// }

// if
// func main() {
// 	nilai := 90

// 	fmt.Println(nilai)
// 	if nilai%2 == 0 {
// 		fmt.Println("Ini angka genap")
// 	} else {
// 		fmt.Println("Ini Angka Ganjil")
// 	}
// }

// switch
// func main() {
// 	name := "joko matic"

// 	switch name {
// 	case "Asep":
// 		fmt.Println("Mantap Asep")
// 	case "joko matic":
// 		fmt.Println("Mantap joko matic")
// 	case "Supaendi":
// 		fmt.Println("Mantap Supaendi")
// 	case "aowkaowoawk":
// 		fmt.Println("Mantap aowkaowoawk")
// 	default:
// 		fmt.Println("iyadeh")
// 	}
// }

// for
// func main() {
// 	names := [...]string{"fadli", "asep", "programmer", "PM"}

// 	for i := 0; i < len(names); i++ {
// 		fmt.Println(names[i])
// 	}

// 	for index, name := range names {
// 		fmt.Print("index", index, "=", name)
// 	}
// }

// function parameter
// func sayName(firstName string, lastName string) {
// 	fmt.Println("nama depanmu = ", firstName)
// 	fmt.Println("nama belakangmu = ", lastName)
// 	fmt.Println(firstName, lastName)
// }

//	func main() {
//		sayName("asep", "junaedi")
//	}
//
// function parameter return multiple values
// func sayName() (string, string) {
// 	return "asep", "koswara"
// }

// func completeName() (awal, tengah, akhir string) {
// 	awal = "asep"
// 	tengah = "koswara"
// 	akhir = "jadi"

// 	return awal, tengah, akhir
// }

// func main() {
// 	firstName, lastname := sayName()

// 	a, b, c := completeName()
// 	fmt.Println(firstName, lastname)
// 	fmt.Println(a, b, c)
// }

//variadic function
// func sumAll(numbers ...int) int {
// 	total := 0

// 	for _, number := range numbers {
// 		total += number
// 	}

// 	return total
// }

// func main() {
// 	fmt.Println(sumAll(100, 10, 10, 101, 10, 10, 110, 10, 1, 10))

// 	numbers := []int{101, 101, 1, 101, 1}

// 	fmt.Println(sumAll(numbers...))
// }

// function as value
// func getGoodbye(name string) string {
// 	return "Good Bye" + name
// }

// func main() {
// 	goodbye := getGoodbye

//		fmt.Println(goodbye(" aokwaokwoakw"))
//	}
//

// function as parameter
// type Filter func(string) string

// func sayHelloWithFilter(name string, filter Filter) {
// 	filteredName := filter(name)

// 	fmt.Println("Hello", filteredName)
// }

// func spamFiltered(name string) string {
// 	if name == "Anjing" {
// 		return "...."
// 	} else {
// 		return name
// 	}
// }

// func main() {

// 	sayHelloWithFilter("Anjing", spamFiltered)

// }

// anonymous function
// type BlackList func(string) bool

// func registerName(name string, blacklist BlackList) {
// 	if blacklist(name) {
// 		fmt.Println("keblok", name)
// 	} else {
// 		fmt.Println("lanjut", name)
// 	}
// }

// func main() {
// 	blacklist := func(name string) bool {
// 		return name == "anjing"
// 	}
// 	registerName("eko", blacklist)
// 	registerName("anjing", blacklist)
// }

// recursive function
// func factorialLoop(value int) int {
// 	// result := 1

// 	// for i := value; i > 0; i-- {
// 	// 	result *= value
// 	// }
// 	// return result

// 	if value == 1 {
// 		return 1
// 	} else {
// 		return value * factorialLoop(value-1)
// 	}

// }

// func main() {
// 	fmt.Println(factorialLoop(1))
// }

//closure
// func main() {
// 	counter := 0

// 	increment := func() {
// 		fmt.Println("Increment")
// 		counter++
// 	}

// 	increment()
// 	increment()
// 	increment()

// 	fmt.Println(counter)
// }

//defer
// func logging() {
// 	fmt.Println("Selesai memanggil function")
// }

// func runApplication() {
// 	defer logging()
// 	fmt.Println("berjalan")
// }

// func main() {
// 	runApplication()
// }

// panic and recover
// func ednApp() {
// 	fmt.Println("selesai app")
// 	message := recover()
// 	fmt.Println("terjadi error", message)
// }

// func runApplication(error bool) {
// 	defer ednApp()

// 	if error {
// 		panic("ups error")
// 	}

// }

// func main() {
// 	runApplication(true)
// 	fmt.Println("asep")
// }

// // struct
// type Customer struct {
// 	Name, Address string
// 	Age           int
// }

// // struct method
// func (customer Customer) sayHello(name string) {
// 	fmt.Println("Hello", name, "My name is", customer.Name)
// }

// func main() {
// 	var asep Customer

// 	fmt.Println(asep)

// 	asep.Name = "mbut"
// 	asep.Address = "bandung"
// 	asep.Age = 12

// 	fmt.Println(asep)
// 	fmt.Println(asep.Name)
// 	fmt.Println(asep.Address)
// 	fmt.Println(asep.Age)

// 	joko := Customer{
// 		Name:    "Kopling",
// 		Address: "Jekardah",
// 		Age:     12,
// 	}

// 	fmt.Println(joko)

// 	joko.sayHello("Joko")
// }

// interface
// type HasName interface {
// 	GetName() string
// }

// type Person struct {
// 	Name string
// }

// type Animal struct {
// 	Name string
// }

// func (person Person) GetName() string {
// 	return person.Name
// }

// func (animal Animal) GetName() string {
// 	return animal.Name
// }

// func SayHello(value HasName) {
// 	fmt.Println("Hello", value.GetName())
// }

// func main() {
// 	person := Person{
// 		Name: "Asep",
// 	}
// 	SayHello(person)

// 	animal := Animal{
// 		Name: "Kambing",
// 	}
// 	SayHello(animal)
// }

//any interface kosong
// func Ups() any {
// 	return "ups"
// }

// func main() {
// 	var kosong any = Ups()

// 	fmt.Println(kosong)
// }

// Nil, hanya bisa digunakan pada function, interface, map,dll gw lupa bjir
// func NewMap(name string) map[string]string {
// 	if name == "" {
// 		return nil
// 	} else {
// 		return map[string]string{
// 			"name": name,
// 		}
// 	}
// }

// func main() {
// 	data := NewMap("oakwaowkoakw")
// 	if data == nil {
// 		fmt.Print("data map masih kosong")
// 	} else {
// 		fmt.Println(data["name"])
// 	}
// }

// type assertions
// func random() any {
// 	return false
// }

// func main() {
// 	var result any = random()
// 	// var resultString string = result.(string)
// 	// fmt.Println(resultString)

// 	switch value := result.(type) {
// 	case string:
// 		fmt.Println("String", value)
// 	case int:
// 		fmt.Println("Integer", value)
// 	default:
// 		fmt.Println("Unknown", value)
// 	}
// }

//POINTER/reference dan operator asterisk -> Kata Kang Eko ini penting

// type Address struct {
// 	City, Province, Country string
// }

// func main() {
// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
// address2 := address1

// address2.City = "Bandung"
// fmt.Println(address1)
// fmt.Println(address2)
// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
// address2 := &address1

// address2.City = "Bandung"
// fmt.Println(address1)
// fmt.Println(address2)

// *address2 = Address{"Jekardah", "DKI", "Indonesia"}
// fmt.Println(address1)
// fmt.Println(address2)

// }

// new, rata-rata pada pake new daripada &
// type Address struct {
// 	City, Province, Country string
// }

// func main() {
// 	var alamat1 *Address = new(Address)
// 	var alamat2 *Address = alamat1

// 	alamat2.Country = "Indonesia"

// 	fmt.Println(alamat1)
// 	fmt.Println(alamat2)
// }

// Pointer di Function
// type Address struct {
// 	City, Province, Country string
// }

// func ChangeAddressToIndonesia(address *Address) {
// 	address.Country = "Indonesia"
// }

// func main() {
// 	address := Address{}
// 	ChangeAddressToIndonesia(&address)

// 	fmt.Println(address)
// }

// pointer method
// type Man struct {
// 	Name string
// }

// func (man *Man) Married() {
// 	man.Name = "Mr. " + man.Name
// }

// func main() {
// 	asep := Man{"oakowakoawk"}
// 	asep.Married()

// 	fmt.Println(asep.Name)
// }

//error

type validationError struct {
	Message string
}
type notFoundError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}
func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation Error"}
	}

	if id != "asep" {
		return &notFoundError{"Data not found"}
	}

	return nil
}

func main() {
	err := SaveData("asep", nil)
	if err != nil {
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("Validation Error: ", validationErr.Error())
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println("not found error: ", notFoundErr.Error())
		// } else {
		// 	fmt.Println("Unknown error: ", err.Error())
		// }

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation error: ", finalError.Error())
		case *notFoundError:
			fmt.Println("Not FOund error: ", finalError.Error())
		default:
			fmt.Println("Unknown error: ", finalError.Error())
		}

	} else {
		fmt.Println("sukses")
	}
}
