package max

import "fmt"

func max(s []int) int {
	l := len(s)
	if l == 2 {
		if s[0] > s[1] {
			return s[0]
		} else {
			return s[1]
		}
	}
	subMax := max(s[1:])
	if s[0] > subMax {
		return s[0]
	} else {
		return subMax
	}
}

func main() {
	r := max([]int{238, 22, 1, 4})
	fmt.Println(r)
}
