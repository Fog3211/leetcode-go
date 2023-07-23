package leetcode

func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)-1

	if letters[right] <= target {
		return letters[0]
	}

	for left < right {
		mid := left + (right-left)>>1
		if letters[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return letters[left]
}
