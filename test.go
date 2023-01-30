package main

func sumAllElelms(a ...int) int {
	res := 0

	for _, e := range a {
		res += e
	}

	return res
}
