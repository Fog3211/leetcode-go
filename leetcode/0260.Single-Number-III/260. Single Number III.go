package leetcode

func singleNumberIII(nums []int) []int {
	eor := 0
	for _, v := range nums {
		eor ^= v
	}
	var rightOne = eor & -eor
	v1 := 0
	v2 := 0
	for _, v := range nums {
		if v&rightOne == 0 {
			v1 ^= v
		} else {
			v2 ^= v
		}
	}
	return []int{v1, v2}
}
