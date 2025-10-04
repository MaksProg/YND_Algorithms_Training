package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, v0, v1, v2 int
	fmt.Scan(&a, &b, &c, &v0, &v1, &v2)

	dist := [3][3]float64{
		{0, float64(a), float64(b)},
		{float64(a), 0, float64(c)},
		{float64(b), float64(c), 0},
	}

	for k := 0; k < 3; k++ {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
	v0f := float64(v0)
	v1f := float64(v1)
	v2f := float64(v2)

	route1 := dist[0][1]/v0f + dist[1][2]/v1f + dist[2][0]/v2f
	route2 := dist[0][2]/v0f + dist[2][1]/v1f + dist[1][0]/v2f
	route3 := dist[0][1]/v0f + dist[1][0]/v1f + dist[0][2]/v0f + dist[2][0]/v1f
	route4 := dist[0][2]/v0f + dist[2][0]/v1f + dist[0][1]/v0f + dist[1][0]/v1f

	minTime := math.Min(math.Min(route1, route2), math.Min(route3, route4))

	fmt.Println(minTime)
}
