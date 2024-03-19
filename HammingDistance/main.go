package main

func main() {

}

func hammingDistance(x int, y int) int {
	ans := 0

	for x != 0 || y != 0 {
		if x&1 != y&1 {
			ans += 1
		}

		x >>= 1
		y >>= 1
	}

	return ans
}
