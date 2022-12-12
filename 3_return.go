package main

func main() {
	var x int = 10
	var result bool = validation(x)
	println(result)
}

func validation(value int) bool {
	if value <= 10 && value >= 20 {
		return true
	} else {
		return false
	}
}
