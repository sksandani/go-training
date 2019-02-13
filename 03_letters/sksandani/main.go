package main

func main() {
}

func letters(s string) map[rune]int {
	resultMap := make(map[rune]int)

	for _, c := range s {
		resultMap[c]++
	}

	return resultMap
}
