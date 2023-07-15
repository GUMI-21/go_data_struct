package algorithm

// F 尾递归斐波那契数
func F(n int, a1, a2 int) int {
	if n == 0 {
		return a1
	}
	return F(n-1, a2, a1+a2)
}

// fib 斐波那契数
func fib(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
