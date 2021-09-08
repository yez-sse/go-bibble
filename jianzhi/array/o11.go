package array

func minArray(numbers []int) int {
	low := 0
	heigh := len(numbers) - 1
	for low < heigh {
		pivot := low + (heigh-low)/2
		if numbers[pivot] < numbers[heigh] {
			heigh = pivot
		} else if numbers[pivot] > numbers[heigh] {
			low = pivot + 1
		} else {
			heigh--
		}
	}
	return numbers[low]

	//for i := 0; i < len(numbers) - 1; i++ {
	//	if numbers[i] > numbers[i + 1] {
	//		return numbers[i + 1]
	//	}
	//}
	//return numbers[0]
}
