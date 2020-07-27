package src

import (
	"math"
)

var rooms = 0
var runnable = 0
var quantity = 0

func Solution(n, d, b int, a []int) int {
	rooms, runnable, quantity = n, d, b
	totalForSide := (rooms + 1) / 2 * quantity
	sum := 0
	enoughElement := 0
	excess := 0
	for i := 0; i < rooms; i++ {
		sum += a[i]
		if sum >= totalForSide {
			excess = sum - totalForSide
			enoughElement = i
			break
		}
	}
	left := make([]int, rooms)
	for i := 0; i < enoughElement; i++ {
		left[i] = a[i]
	}
	left[enoughElement] = a[enoughElement] - excess
	security1 := check(left, (rooms+1)/2)

	right := make([]int, rooms)
	for i := 0; i < rooms-enoughElement; i++ {
		right[i] = a[rooms-1-i]
	}
	right[rooms-1-enoughElement] = excess
	security2 := check(right, rooms/2)
	return int(math.Max(security1, security2))
}

func check(values []int, middle int) float64 {
	val := 0
	res := 0
	for i := 0; i < middle; i++ {
		run := float64((runnable+1)*i + runnable)
		for j := (runnable + 1) * i; float64(j) <= math.Min(float64(rooms-1), run); j++ {
			val += values[j]
		}
		if val >= quantity {
			val -= quantity
		} else {
			res++
		}
	}
	return float64(res)
}
