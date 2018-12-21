package slices

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	fmt.Println("s1=", s1)
	s2 := s1[3:5]
	fmt.Println("s2 = ", s2)
	fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d\n", s2, len(s2), cap(s2))
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)

	fmt.Println("s3,s4,s5=", s3, s4, s5)
	fmt.Println("cap(arr),arr = ", cap(arr), arr)

}
