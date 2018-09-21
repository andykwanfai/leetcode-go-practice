func longestCommonPrefix(strs []string) string {
	short := shortest(strs)

	//compare each char
	for i := range short {
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != short[i] {
				return strs[j][:i]
			}
		}
	}

	return short
}

//find the shortest string
func shortest(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]
	for _, s := range strs {
		if len(res) > len(s) {
			res = s
		}
	}

	return res
}