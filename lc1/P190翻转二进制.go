package lc1

func reverseBits(num uint32) uint32 {
	var res uint32 = 0
	for i := 0; i < 32 && num != 0; i++ {
		res |= (num & 1) << (31 - i)
		num >>= 1
	}
	return res
}

func hammingWeight(num uint32) int {
	res := 0
	for i := 0; i < 32 && num != 0; i++ {
		if (num & 1) == 1 {
			res++
		}
		num >>= 1
	}
	return res
}
