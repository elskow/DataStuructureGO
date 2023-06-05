package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m + n - 1; i >= 0; i-- {
		if m == 0 {
			nums1[i] = nums2[n-1]
			n--
		} else if n == 0 {
			nums1[i] = nums1[m-1]
			m--
		} else if nums1[m-1] > nums2[n-1] {
			nums1[i] = nums1[m-1]
			m--
		} else {
			nums1[i] = nums2[n-1]
			n--
		}
	}
}

func main() {
	nums := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	merge(nums, m, nums2, n)
	for _, num := range nums {
		println(num)
	}
}
