import (
	"math"
)

func reverse(x int) int {
	reverse := 0
	for x != 0 {
		reverse = reverse*10 + x%10
		x /= 10
	}

	if reverse > math.MaxInt32 || reverse < math.MinInt32 {
		reverse = 0
	}
	return reverse
}