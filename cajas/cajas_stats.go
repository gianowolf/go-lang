package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	tries int = 50
)

var (
	// boxes  = []int{1, 2, 0, 4, 5, 3}
	boxes  []int
	values []bool
)

func newRand() *rand.Rand {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r
}

func orderedSliceMake(slice_size int) []int {

	s := make([]int, slice_size, slice_size)

	for i := 0; i < slice_size; i++ {
		s[i] = i
	}

	return s
}

func fillBoxes(size int) []int {

	r := newRand()

	disordered_boxes := make([]int, 0, size)
	ordered_boxes := orderedSliceMake(size)

	for i := size; i > 0; i-- {
		pos := r.Intn(i)
		num := ordered_boxes[pos]
		disordered_boxes = append(disordered_boxes, num)
		ordered_boxes = append(ordered_boxes[:pos], ordered_boxes[pos+1:]...)
	}

	return disordered_boxes
}

func escape(size int) int {

	boxes := fillBoxes(size)

	for prisioner := 0; prisioner < size; prisioner++ {
		success := false
		max_attemps := size / 2
		attemp := 1

		paper := boxes[prisioner]
		next_box := paper

		for {
			attemp++
			paper = boxes[next_box]

			if prisioner == paper {
				success = true
				break
			}

			next_box = paper

			if attemp >= max_attemps {
				break
			}
		}

		if !success {
			return 0
		}
	}
	return 1
}

func newExperiment(size int) float64 {

	var success int
	for i := 0; i < tries; i++ {
		success += escape(size)
	}

	return float64(success) / float64(tries)
}

func main() {

	size := 300_000
	fmt.Printf("promedio de exito con %d cajas: %f\n", size, newExperiment(size))
	// for size < 200_000 {
	// 	fmt.Printf("promedio de exito con %d cajas: %f\n", size, newExperiment(size))
	// 	size *= 10
	// }

}
