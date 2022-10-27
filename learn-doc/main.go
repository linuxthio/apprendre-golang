package main

import (
	"fmt"
	"os"
)

func main() {
	n, er := fmt.Println("vim-go")
	fmt.Println(n, er)

	os.Remove("hello")
}
