package main
import "fmt"
type Vertex struct {
	X, Y int
}

func main() {
	v := new(Vertex)
	fmt.Println(v)
	v.X, v.Y = 9, 5
	fmt.Println(v)
}
