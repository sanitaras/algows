package main

import (
	"fmt"
	"net/http"
	"reflect"
)

func main() {

	a5 := []int{0, 1, 1, 2, 3}
	a10 := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	a15 := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377}

	b5 := getFibs(5)
	b10 := getFibs(10)
	b15 := getFibs(15)

	z5 := reflect.DeepEqual(a5, b5)
	z10 := reflect.DeepEqual(a10, b10)
	z15 := reflect.DeepEqual(a15, b15)

	if z5 && z10 && z15 {
		fmt.Println("Test Fibonacci Sequence: PASS")
		http.HandleFunc("/", getHandler)
		http.HandleFunc("/fib/", getHandler)
		http.HandleFunc("/alg1/", getHandler)
		http.HandleFunc("/alg2/", getHandler)
		http.HandleFunc("/alg3/", getHandler)
		fmt.Println("Algorithms RESTful service started.")
		http.ListenAndServe(":9000", Log(http.DefaultServeMux))
	} else {
		fmt.Println("Test Fibonacci Sequence: FAIL")
	}
}
