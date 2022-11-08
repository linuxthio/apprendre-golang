package polymorphe

import "fmt"


func area(w,h int) int {
	return w*h
}
func area(w,h float64) float64 {
	return w*h
}

func main(){
	fmt.Println("")
}