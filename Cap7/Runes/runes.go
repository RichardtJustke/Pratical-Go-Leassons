package main

import (
"fmt"
)


func main() {
	s := "我愛 golang"
	for _, v := range s {// o v é uma runa
		fmt.Printf("Unicode code point: %U - character ´%c´ - binary %b -hex %x - decimal %d/n", v, v, v, v, v)
	}
} 
