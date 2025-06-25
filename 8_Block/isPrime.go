package main

import "fmt"

func main() {
	const limit = 100
	primes := findPrimes(limit)
	fmt.Printf("Простые числа до %d:\n%v\n", limit, primes)
}

func findPrimes(limit int) []int {

	sieve := make([]bool, limit+1)

	sieve[0], sieve[1] = true, true

	for i := 2; i*i <= limit; i++ {
		if !sieve[i] {
			for j := i * i; j <= limit; j += i {
				sieve[j] = true
			}
		}
	}

	var primes []int
	for num, isComposite := range sieve {
		if !isComposite {
			primes = append(primes, num)
		}
	}

	return primes
}
