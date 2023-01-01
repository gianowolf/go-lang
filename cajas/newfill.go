package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	size int = 8
)

var (
	// boxes  = []int{1, 2, 0, 4, 5, 3}
	boxes  []int
	values []bool
)

func fillBoxes() []int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	disordered_boxes := make([]int, 0, size)
	ordered_boxes := make([]int, size, size)

	for i := 0; i < size; i++ {
		ordered_boxes[i] = i
	}

	fmt.Println("Ordenadas:", ordered_boxes)
	fmt.Println("Desordenadas:", disordered_boxes)
	fmt.Println()

	for i := size; i > 0; i-- {
		fmt.Println("Iteracion: ", i)

		pos := r.Intn(i)
		fmt.Println("Posicion:", pos)

		num := ordered_boxes[pos]
		fmt.Println("Numero: ", num)

		disordered_boxes = append(disordered_boxes, num)
		ordered_boxes = append(ordered_boxes[:pos], ordered_boxes[pos+1:]...)

		fmt.Println("Ordenadas:", ordered_boxes)
		fmt.Println("Desordenadas:", disordered_boxes)
		fmt.Println()
	}

	return disordered_boxes

}

func main() {

	boxes := fillBoxes()
	fmt.Println(boxes)

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
