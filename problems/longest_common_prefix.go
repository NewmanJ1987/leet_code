package problems

func LongestCommonPrefix(strs []string) string {
	shortest_string := shortestString(strs)
	longest_substring := ""

	for i := 1; i <= shortest_string; i++ {
		substring_arr := substringArray(strs, i)
		longest_substring_c := matchSubstring(substring_arr, longest_substring)
		if len(longest_substring) < len(longest_substring_c) {
			longest_substring = longest_substring_c
		} else {
			break
		}
	}

	return longest_substring
}

func shortestString(strs []string) int {
	shortest := len(strs[0])
	for _, value := range strs {
		if len(value) < shortest {
			shortest = len(value)
		}
	}
	return shortest
}

func substringArray(strs []string, size int) []string {
	substring_arr := make([]string, 0)
	for _, value := range strs {
		substring_arr = append(substring_arr, value[:size])
	}
	return substring_arr
}

func matchSubstring(substring_arr []string, previous_longest string) string {
	match := true
	for i := 0; i < len(substring_arr)-1; i++ {
		if substring_arr[i] != substring_arr[i+1] {
			match = false
			break
		}
	}
	if match {
		return substring_arr[0]
	}
	return previous_longest
}
