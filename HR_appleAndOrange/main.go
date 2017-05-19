package main

import "fmt"

func main() {
	var s, t, a, b, napples, noranges, d, fallsapples, fallsoranges int

	fmt.Scanf("%d %d", &s, &t)
	fmt.Scanf("%d %d", &a, &b)
	fmt.Scanf("%d %d", &napples, &noranges)

	for i := 0; i < napples; i++ {
		fmt.Scanf("%d", &d)
		var dist = a + d
		if dist >= s && dist <= t {
			fallsapples++
		}
	}
	for i := 0; i < noranges; i++ {
		fmt.Scanf("%d", &d)
		var dist = b + d
		if dist >= s && dist <= t {
			fallsoranges++
		}
	}

	fmt.Println(fallsapples)
	fmt.Println(fallsoranges)
}
