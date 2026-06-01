func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := len(nums1) - 1,len(nums2)-1,m-1

	for i >= 0 {
		if j < 0 {
			nums1[i] = nums1[k]
			k--
		}else if k < 0{
			nums1[i] = nums2[j]
			j--
		}else{
			if nums1[k] < nums2[j]{
				nums1[i] = nums2[j]
				j--
			}else{
				nums1[i] = nums1[k]
				k--
			}
		}
		i--
	}
}
