package main

import "fmt"

// test29 input 1571 4240 9023 4234 	-> YES
// test10 input 43 2 70 2			-> NO

func main() {
	var x1, v1, x2, v2 float32
	fmt.Scanf("%f %f %f %f", &x1, &v1, &x2, &v2)

	if x1 > x2 {
		x1, v1, x2, v2 = x2, v2, x1, v1
	}

	if x1 == x2 {
		fmt.Println("YES")
		return
	}

	if (x2-x1) < (x2+v2-x1-v1) || v1 == v2 {
		fmt.Println("NO")
		return
	}

	for {
		if x1 > x2 {
			fmt.Println("NO")
			return
		}

		if x1 == x2 {
			fmt.Println("YES")
			return
		}

		x1 = x1 + v1
		x2 = x2 + v2
	}
}
