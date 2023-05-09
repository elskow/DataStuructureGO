package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var totalLen = len(nums1) + len(nums2)
	if totalLen%2 == 0 {
		return float64(findKth(nums1, nums2, totalLen/2)+findKth(nums1, nums2, totalLen/2+1)) / 2.0
	} else {
		return float64(findKth(nums1, nums2, totalLen/2+1))
	}
}

func findKth(nums1 []int, nums2 []int, k int) int {
	var len1 = len(nums1)
	var len2 = len(nums2)
	if len1 > len2 {
		return findKth(nums2, nums1, k)
	}
	if len1 == 0 {
		return nums2[k-1]
	}
	if k == 1 {
		return min(nums1[0], nums2[0])
	}
	var i = min(len1, k/2)
	var j = min(len2, k/2)
	if nums1[i-1] > nums2[j-1] {
		return findKth(nums1, nums2[j:], k-j)
	} else {
		return findKth(nums1[i:], nums2, k-i)
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
