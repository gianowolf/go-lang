package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	type product struct {
		ID                 string   `json:"id"`
		Title              string   `json:"title"`
		Description        string   `json:"description"`
		Price              string   `json:"price"`
		DiscountPercentage string   `json:"discountPercentage"`
		Rating             string   `json:"rating"`
		Stock              string   `json:"stock"`
		Brand              string   `json:"brand"`
		Category           string   `json:"category"`
		Thumbnail          string   `json:"thumbnail"`
		Images             []string `json:"images"`
	}

	resp, err := http.Get("https://dummyjson.com/products/category/smartphones")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
