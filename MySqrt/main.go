package mysqrt

func main() {

}

func mySqrt(x int) int {

	if x < 2 {
		return x
	}

	l, r := 1, x

	for l <= r {
		middle := (l + r) / 2

		pow := middle * middle

		if pow == x {
			return middle
		}

		if pow > x {
			r = middle - 1
		} else {
			l = middle + 1
		}
	}

	return r
}
