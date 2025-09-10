package piscine

func Index(s string, toFind string) int {
	n := len(toFind)
	if n == 0 {
		return 0
	}
	for i := 0; i <= len(s)-n; i++ {
		if s[i:i+n] == toFind {
			return i
		}
	}
	return -1
}
