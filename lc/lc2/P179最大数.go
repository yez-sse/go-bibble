package lc2

import (
	"sort"
	"strconv"
	"strings"
)

func largestNum(nums []int) string {
	ss := make([]string, len(nums))
	for i, num := range nums {
		ss[i] = strconv.Itoa(num)
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i]+ss[j] >= ss[j]+ss[i]
	})
	res := strings.Join(ss, "")
	if res[0] == '0' {
		return "0"
	}
	return res
}
