package init

import "fmt"

type InitGo struct {
}

func Test() {
	fmt.Println(" ")
}

func init() {
	fmt.Println("i am init func")
}
