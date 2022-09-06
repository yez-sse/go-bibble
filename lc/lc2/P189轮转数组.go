package lc2

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-k-1)
	reverse(nums, len(nums)-k, len(nums)-1)
	reverse(nums, 0, len(nums)-1)
}

func reverse(nums []int, left, right int) {
	for left < right {
		temp := nums[left]
		nums[left] = nums[right]
		nums[right] = temp
		left++
		right--
	}
}
