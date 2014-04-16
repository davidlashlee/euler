/*
Problem 12: Highly Divisble Triangular Number
The sequence of triangle numbers is generated by adding the natural numbers. So the 7th triangle number would be 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28.
The first ten terms would be: 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...
Let us list the factors of the first seven triangle numbers:
 1: 1
 3: 1,3
 6: 1,2,3,6
10: 1,2,5,10
15: 1,3,5,15
21: 1,3,7,21
28: 1,2,4,7,14,28
We can see that 28 is the first triangle number to have over five divisors.
What is the value of the first triangle number to have over five hundred divisors?
Answer: 76576500
*/

package problems

import "github.com/zolrath/euler/util/primes"

const ANSWER_012 = 76576500

func totalDivisors(n int) int {
	pfs := primes.PrimeFactors(n)
	total := 1
	for _, e := range pfs {
		total = (e + 1) * total
	}
	return total
}

func Euler012() int {
	var tnum, divisors int
	for i := 1; divisors <= 500; i++ {
		tnum += i
		divisors = totalDivisors(tnum)
	}
	return tnum
}

// func primeFactors(n int) map[int]int {
// 	pfs := map[int]int{}
// 	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
// 		for n%i == 0 {
// 			n = n / i
// 			pfs[i] = pfs[i] + 1
// 		}
// 	}
// 	if n > 1 {
// 		pfs[n] = 1
// 	}
// 	return pfs
// }
