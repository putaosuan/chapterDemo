package testcase1

func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}
func getSub(n, m int) int {
	return n - m
}
