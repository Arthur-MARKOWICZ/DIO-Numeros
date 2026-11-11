package main

import "fmt"

func main() {
	var i int16 = 1
	for i <= 100 {
		if i%3 == 0 {
			fmt.Println(i)
		}
		i++
	}

}
