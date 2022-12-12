package main

import "fmt"

func f1() {
	fmt.Println("F1: Hello World")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s \n", p1, p2)
}

func f3() string {
	return "F3: Hello World"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "F5: Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("Hello", "World")
	fmt.Println(f3())
	fmt.Println(f4("Hello", "World"))
	fmt.Println(f5())
}
