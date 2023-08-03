package main

import "fmt"

func calculateFibonacci(n int) []int {
	var fibo []int
	for i := 0; i < n; i++ {
		if i < 2 {
			fibo = append(fibo, i)
		} else {
			fibo = append(fibo, fibo[i-1]+fibo[i-2])
		}
	}

	return fibo
}

func main() {
	fmt.Println("Fibonacci Series")
	fmt.Println("1st Fibonacci Number is", calculateFibonacci(1))
	fmt.Println("2nd Fibonacci Number is", calculateFibonacci(2))
	fmt.Println("3rd Fibonacci Number is", calculateFibonacci(3))
	fmt.Println("4th Fibonacci Number is", calculateFibonacci(20))
}
