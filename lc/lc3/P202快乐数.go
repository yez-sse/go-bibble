package lc3

func isHappy(n int) bool {
	slow := n
	fast := bitSquareSum(n) //快的追上慢的时候，是可能多跑了n圈，而不是说一起开始的时候，就刚好多一圈，也是多n圈的
	for fast != 1 && slow != fast {
		slow = bitSquareSum(slow)
		fast = bitSquareSum(bitSquareSum(fast))
	}
	return fast == 1 //这里用slow会有计算不对的，要仔细判别与java写法的不同
}

func bitSquareSum(n int) int {
	sum := 0
	for n > 0 {
		temp := n % 10
		sum += temp * temp
		n /= 10
	}
	return sum
}
