package main

import "fmt"
import "math"

func main() {
	var x1, x2, v1, v2 float64
	fmt.Scanf("%f %f %f %f", &x1, &v1, &x2, &v2)

	var oldDistance, newDistance = 0.0, math.Abs(x1 - x2)

	for {
		oldDistance = newDistance
		newDistance = math.Abs(x1 - x2)

		if oldDistance < newDistance {
			fmt.Println("NO")
			return
		}

		if newDistance == 0 {
			fmt.Println("YES")
			return
		}

		x1 = x1 + v1
		x2 = x2 + v2
	}
}
