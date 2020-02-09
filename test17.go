package main

func main() {
	nums1 := []int{2, 0}
	nums2 := []int{1}
	merge(nums1, 1, nums2, 1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		for index, value := range nums2 {
			nums1[index] = value
		}
		return
	}
	for index, value := range nums2 {
		i := m + index - 1
		for value < nums1[i] {
			nums1[i+1] = nums1[i]
			i--
			if i < 0 {
				break
			}
		}
		nums1[i+1] = value
	}
}
