package main

import (
	"fmt"
	"github.com/mattn/go-precure"
)

func main() {
	for _, a := range precure.AllStars {
		fmt.Println(a.HumanName())
	}
}
