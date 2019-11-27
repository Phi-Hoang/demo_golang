package main

import (
	"fmt"
	"strconv"
)

/* type Student struct {
	FirstName string `json:"first_name" bson:"full_name" validate:"required"`
	LastName  string
	Age       int
	Email     string
}

func (s Student) GetFullName() string {
	return s.FirstName + " " + s.LastName
} */

/* func (s Student) SetEmail(email string) {
	s.Email = email
} */
func (s *Student) SetEmail(email string) {
	s.Email = email
}

func main() {
	fmt.Println("Hello")
	// Int
	aInt := 1
	fmt.Println(aInt)

	var bInt int
	bInt = 2
	fmt.Printf("bInt = %d \n", bInt)

	// String
	aString := "Hello S"
	fmt.Printf("aString = %s \n", aString)

	// Array/Slide
	aSlice := []string{"a", "b", "c", "d"}
	fmt.Println(aSlice)

	bSlice := []int{2, 3, 4, 6, 2, 8, 3, 7, 5, 9}
	fmt.Printf("bSlice = %v \n", bSlice)
	for i, v := range bSlice {
		fmt.Printf("%d %d\n", i, v)
	}
	for _, v := range bSlice {
		if v > 5 {
			fmt.Printf("%d ", v)
		}
	}

	aArray := [2]int{1, 3}
	fmt.Println(aArray)

	// Map
	aMap := map[string]int{"age": 1000, "level": 3}
	fmt.Println(aMap)

	// Struct
	/* type Student struct {
		FirstName string `json:"first_name" bson:"full_name" validate:"required"`
		LastName  string
		Age       int
		Email     string
	} */

	aSt := Student{
		FirstName: "Phi",
		LastName:  "Hoang",
		Age:       30,
		Email:     "phi@email.com",
	}
	bSt := Student{"Hoang", "Phi", 40, "hoang@name.com"}
	fmt.Println(aSt)
	fmt.Printf("bSt %+v \n", bSt)

	// Interface
	var i interface{} // use as "any" type to store unknow type var
	i = aSt
	fmt.Println(i)

	// Channel
	ch := make(chan int, 2)
	ch <- 1           // write into channel
	ch <- 12          // write into channel
	fmt.Println(<-ch) // read from channel
	fmt.Println(<-ch) // read from channel

	// Function
	//SetEmail(aSt, "abc")
	SetEmail(&aSt, "bbbb")
	fmt.Println(aSt)

	// Function Receiver
	fmt.Println(aSt.GetFullName())
	aSt.SetEmail("aaaa")
	fmt.Println(aSt)

	// int <-> string
	cInt := 100
	cString1 := strconv.Itoa(cInt)
	fmt.Println(cString1)
	cString2 := fmt.Sprintf("%d", cInt)
	fmt.Println(cString2)
	//cInt2 := strconv.ParseInt(cString2)
	//fmt.Printf("%d \n", cInt2)

	// parse inputJson to struct
	inputJson := `[
		{"first_name": "Victor", "last_name": "Nguyen", "age": 100, "class_name":"golang"},
		{"first_name": "Anh", "last_name": "Dinh", "age":200, "class_name":"golang"}
		]`
	fmt.Println(inputJson)
	// encoding/json
}

/* func SetEmail(s Student, email string) {
	s.Email = email
} */
func SetEmail(s *Student, email string) {
	s.Email = email
}
