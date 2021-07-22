package main

import "fmt"

func minutesToSeconds(m int) int {
	return m * 60
}

func main() {
	var minutes int
	fmt.Printf("Minutos: ")
	fmt.Scanf("%d\n", &minutes)
	seconds := minutesToSeconds(minutes)
	fmt.Printf("%d minutos é o mesmo que %d segundos ", minutes, seconds)
}
