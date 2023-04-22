package main

func main() {
	array := [4]int{19, 86, 1, 12}
	sum := 0

	for _, value := range array {
		sum += value
	}
	println(sum)
}
