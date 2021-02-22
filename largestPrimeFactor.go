package main

import "fmt"

// The number was specified in the Euler problem
// Delete this and add just the number for a generalized solution
const NUMBER = 600851475143

func main() {

	fmt.Println(PrimeFactors(NUMBER))

}

// PrimeFactors gets all prime factors of  given nmumber n
func PrimeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point, so we can skip one element
	// note: i = i + 2
	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}
	return
}