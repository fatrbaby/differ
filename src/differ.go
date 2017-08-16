package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("\n %c[1;30;41m%s%c[0m\n\n", 0x1B, "input a path you want to scan.", 0x1B)
		os.Exit(0)
	}

	for _, pathName := range os.Args[1:]{
		fmt.Println(pathName)
	}
}
