package array

func fib(n int) int {
	f0, f1 := 0, 1
	for i := 0; i < n; i++ {
		f0, f1 = f1, (f0+f1)%1000000007
	}
	return f0

	if n == 0 || n == 1 {
		return n
	}
	res := make([]int, n+1) // 如何定义指定长度数组并初始化
	res[0] = 0
	res[1] = 1
	for i := 2; i <= n; i++ {
		res[i] = (res[i-1] + res[i-2]) % 1000000007
	}
	return res[n]
}
