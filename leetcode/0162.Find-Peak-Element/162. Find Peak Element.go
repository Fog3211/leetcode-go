package leetcode

func findPeakElement(nums []int) int {
	if len(nums) <= 1 || nums[0] > nums[1] {
		return 0
	} else if nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	} else {
		left, right := 0, len(nums)-1

		for left < right {
			mid := left + (right-left)>>1

			if nums[mid] > nums[mid+1] {
				right = mid
			} else {
				left = mid + 1
			}
		}

		return left
	}
}
