package array_product

func ProductExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	for i:=0; i<len(res)-1; i++ {
		res[i+1] = res[i]*nums[i]
	}

	prec := 1
	for i:=len(nums)-1; i>0; i-- {
		prec = prec * nums[i]
		res[i-1] = res[i-1] * prec
	}

	return res
}