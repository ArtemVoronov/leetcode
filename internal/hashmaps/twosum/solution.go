package twosum

func twoSum(nums []int, target int) []int {
	l := len(nums)
	for i := 1; i < l; i++ {
		for j := i; j < l; j++ {
			// fmt.Printf("i: %v, j: %v, nums[j]: %v, nums[j-i]: %v\n", i, j, nums[j], nums[j-i])
			if nums[j]+nums[j-i] == target {
				return []int{j - i, j}
			}
		}
	}

	return []int{}
}
