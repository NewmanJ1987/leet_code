package problems

func MyAlgorithm(m map[string]int, input string) int {
	value := 0
	calculate_last := true
	for i := 0; i < len(input); i++ {
		if i < len(input)-1 {
			curr := string(input[i])
			next := string(input[i+1])
			if m[curr] >= m[next] {
				value += m[curr]
			} else {
				value += m[next] - m[curr]
				i++
				if i+1 == len(input) {
					calculate_last = false
				}
			}

		}
	}
	if calculate_last {
		last := string(input[len(input)-1])
		value += m[last]
	}
	return value
}

func EfficentAlgorithm(m map[string]int, input string) int {
	value := 0
	for i := 0; i < len(input); i++ {
		curr := string(input[i])
		if i != 0 {
			last := string(input[i-1])
			if m[curr] > m[last] {
				value -= 2 * m[last]
			}
		}
		value += m[curr]

	}
	return value
}

func EfficentAlgorithmRecursive(m map[string]int, input string, last string, acc int) int {
	curr := string(input[0])
	if m[last] < m[curr] {
		acc -= 2 * m[last]
	}
	if len(input) == 1 {
		return acc + m[input]
	}
	substring := input[1:]
	return EfficentAlgorithmRecursive(m, substring, curr, acc+m[curr])
}
