package cat

func Years(y int) int {
	return y * 7
}

func YearsSum(z int) int {
	count := 0
	for i := 0; i < z; i++ {
		count += 7
	}
	return count
}
