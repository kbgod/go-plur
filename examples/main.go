package main

import (
	"fmt"

	"github.com/kbgod/go-plur"
)

func main() {
	one := "яблоко"
	two := "яблока"
	many := "яблук"

	v := 112
	fmt.Println(plur.NumberText(v, one, two, many))
}
