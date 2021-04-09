package merge_sorted_array

func Merge(nums1 []int, m int, nums2 []int, n int)  {
	var i, j int

	for i < m+n && j < n {
		if (n-j) == (n+m-i) {
			nums1[i] = nums2[j]
			j++
		} else if nums1[i] > nums2[j] {
			addNumAndShift(i, nums2[j], nums1)
			j++
		}
		i++
	}
}

func addNumAndShift(idx, num int, arr []int) {
	for i:=len(arr)-1; i>idx; i-- {
		arr[i] = arr[i-1]
	}
	arr[idx] = num
}
