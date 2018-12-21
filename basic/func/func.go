package _func

import "fmt"

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operator: " + op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func pointTest() {
	var a int = 2
	var b int = 10
	var pa *int = &a
	*pa = 5
	fmt.Println(a)
	pa = &b
	*pa = 100
	fmt.Println(b)
}

func main() {
	fmt.Println(eval(3, 4, "*"))
	q, r := div(13, 3)
	fmt.Println(q, r)

	pointTest()

}
