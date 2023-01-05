package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	boxes []int
)

func newRand() *rand.Rand {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r
}

func orderedSliceMake(slice_size int) []int {

	s := make([]int, slice_size, slice_size)

	for i := 0; i < size; i++ {
		s[i] = i
	}

	return s
}

func fillBoxes() []int {

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
func main() {

	start := time.Now()
	boxes = fillBoxes()
	elapsed := time.Since(start)
	fmt.Println(boxes, elapsed)

	return
	for prisioner := 0; prisioner < size; prisioner++ {

		fmt.Println()
		fmt.Println("Prisionero ", prisioner)
		max_attemps := size / 2
		attemp := 1

		paper := boxes[prisioner]
		fmt.Printf("Intento %d: Encontro el papel %d en la caja con su numero \n", attemp, paper)
		next_box := paper

		for {
			attemp++
			paper = boxes[next_box]
			fmt.Printf("Intento %d: Encontro el papel %d en la caja %d \n", attemp, paper, next_box)

			if prisioner == paper {
				fmt.Println("Encontro el papel con su numero en la caja", next_box)
				break
			}

			next_box = paper

			if attemp >= max_attemps {
				fmt.Println("El prisionero llego su limite de intentos")
				break
			}
		}
	}
}
