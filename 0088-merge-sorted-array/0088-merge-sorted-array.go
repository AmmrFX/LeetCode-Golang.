func merge(nums1 []int, m int, nums2 []int, n int)  {
        	count:=0
 for i := m; i < len(nums1); i++ {
	nums1[i] =nums2[count]
	count++
 }
  sort.Ints(nums1)
}