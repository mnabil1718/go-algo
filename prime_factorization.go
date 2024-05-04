package main

func Factor(primes []int, number int) []int {

	var storage []int

	if len(primes) < 1 {
		return []int{number}
	}

	var i int
	for i < len(primes) {
		if number%primes[i] == 0 {
			number = number / primes[i]
			storage = append(storage, primes[i])
		} else {
			i++
		}
	}

	return storage
}
