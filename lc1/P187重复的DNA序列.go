package lc1

func findRepeatedDnaSequences(s string) []string {
	L := 10
	var res []string
	siMap := make(map[string]int)
	for i := 0; i <= len(s)-L; i++ {
		subStr := s[i : i+L]
		siMap[subStr]++
		if siMap[subStr] == 2 {
			res = append(res, subStr)
		}
	}
	return res
}
