package piscine

func Abort(a, b, c, d, e int) int {
	nums := []int{a, b, c, d, e}

	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums[2]
}
