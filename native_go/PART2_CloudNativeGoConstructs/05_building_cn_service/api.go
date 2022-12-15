package main

import (
	"errors"
	"fmt"
)

var capitals = make(map[string]string)
var ErrorNoSuchKey = errors.New("no such key")

func main() {
	var capital string
	var e error

	fmt.Println(capitals)

	Put("Brasil", "Brasilia")
	Put("Argentina", "Buenos Aires")
	Put("Bolivia", "La Paz")

	fmt.Println(capitals)

	capital, e = Get("Argentina")
	fmt.Println(capital, e)

	capital, e = Get("Brasil")
	fmt.Println(capital, e)
}

func Put(key, value string) error {
	capitals[key] = value
	return nil
}

func Get(key string) (string, error) {
	value, ok := capitals[key]

	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func Delete(key string) error {
	delete(capitals, key)

	return nil
}
