package main

import "fmt"

func main() {
	var colorsTwo map[string]string
	fmt.Println(colorsTwo)

	colorsThree := make(map[int]string)
	colorsThree[0] = "000"
	delete(colorsThree, 0)
	fmt.Println(colorsThree)

	colorsOne := map[string]string{
		"white": "000",
		"black": "fff",
		"gray":  "ccc",
	}

	fmt.Println(colorsOne)

	for color, hex := range colorsOne {
		fmt.Print(color, "-", hex, "\n")
	}

}
