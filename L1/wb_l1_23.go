// удалить i-ый элемент из слайса

package main

import (
	"fmt"
)

// короткий способ
func shortMethod[T any](sl []T, i int) ([]T, T) {
	rem := sl[i]
	sl = append(sl[:i], sl[i+1:]...)
	return sl, rem
}

// быстрый способ
func quickMethod[T any](sl []T, i int) ([]T, T) {
	rem := sl[i]
	sl[i] = sl[len(sl)-1]
	sl = sl[:len(sl)-1]
	return sl, rem
}

//медленные способ
func slowMethod[T any](sl []T, i int) ([]T, T) {
	rem := sl[i]
	copy(sl[i:], sl[i+1:])
	sl = sl[:len(sl)-1]
	return sl, rem
}

func myPrint[T any](sl []T, rem T) {
	fmt.Println(sl)
	fmt.Println("Removed element:", rem)
	fmt.Println()
}

func main() {
	i := 3
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := []string{"A", "B", "C", "D", "E"}
	c := []float32{3.4, 5.67, 76.89, 2.47, 67.1}

	fmt.Println("Quick solution:")
	a, rem := quickMethod(a, i)
	myPrint(a, rem)

	fmt.Println("Slow solution:")
	b, rem1 := slowMethod(b, i)
	myPrint(b, rem1)

	fmt.Println("Short solution:")
	c, rem2 := shortMethod(c, i)
	myPrint(c, rem2)
}
