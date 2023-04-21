package main
import "fmt"
func fib(n int) int {
	const mod int = 1e9 + 7
	if n < 2 {
		return n
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = (p + q) % mod
	}
	return r
}
func main() {
	n := 0
	fmt.Scan(&n)
	res := fib(n)
	fmt.Println(res)
}