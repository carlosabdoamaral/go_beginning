package main

func notaParaConceito(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 7 && n < 9 {
		return "B"
	} else {
		return ""
	}
}

func main() {
	notaParaConceito(9)
}
