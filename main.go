package main

import (
	"errors"
	"fmt"
	"os"
)

// fungsi
func add(a int, b int) int {
	return a + b
}

func swap(x, y string) (string, string) {
	return y, x
}

// struct
type Person struct {
	Name string
	Age  int
}

// interface
type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow"
}

// error handling
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	// menampilkan hello world
	fmt.Println("Hello, World!")

	// variabel dan tipe data
	var x int = 10  // -> tipe data ditentukan manual
	var y = "Hello" // -> tipe data ditentukan otomatis
	z := 25         // -> penugasan singkat
	fmt.Println(x, y, z)

	// fungsi
	fmt.Println(add(3, 4))
	fmt.Println(swap("first", "second"))

	// flow control (conditional statements)
	a := 15
	if a > 10 {
		fmt.Println("Greater than 10")
	} else {
		fmt.Println("Less or equal to 10")
	}

	// flow control (switch statements)
	b := 2
	switch b {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Other")
	}

	// flow control (looping -> for)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// arrays, slices, and maps
	// -> arrays
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// -> slices
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// -> maps
	m := make(map[string]int)
	m["age"] = 30
	fmt.Println(m["age"])

	// struct
	p := Person{Name: "John", Age: 30}
	fmt.Println(p)

	// concurrency dengan goroutines dan channels
	// -> goroutines
	go func() {
		fmt.Println("Hello from goroutine")
	}()
	fmt.Println("Hello from main")

	// -> channels
	ch := make(chan int)

	go func() {
		ch <- 5
	}()

	fmt.Println(<-ch)

	// interface
	var animal Animal

	animal = Dog{}
	fmt.Println(animal.Speak())

	animal = Cat{}
	fmt.Println(animal.Speak())

	// error handling
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// pustaka standar
	fmt.Println("Hello, World!")
	fmt.Println("Current Directory:", os.Getenv("PWD"))
}
