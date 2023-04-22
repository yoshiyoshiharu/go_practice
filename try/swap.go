package main

func main() {
	n, m := swap(10, 20)
	println(n, m)
}

func swap(n int, m int) (int, int) {
  return m, n
}
