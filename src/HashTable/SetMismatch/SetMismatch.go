package SetMismatch

//func findErrorNums(nums []int) []int {
//	sameNum := 0
//	temp := 0
//	sum := 0
//	sort.Sort(sort.IntSlice(nums))
//	for _,value := range nums {
//		if temp == 0 {
//			temp = value
//		}else{
//			if temp == value {
//				sameNum = value
//			}
//			temp = value
//		}
//		sum += value
//
//	}
//	return []int{sameNum,((1+len(nums))*len(nums)/2)-(sum-sameNum)}
//}

//func findErrorNums(nums []int) []int {
//	type void struct{}
//	var value void
//	set := make(map[int]void)
//
//	sameNum := 0
//	length := len(nums)
//	sum := 0
//
//	for i:=0;i<length;i++ {
//		if _,ok:=set[nums[i]];ok {
//			sameNum = nums[i]
//		}else{
//			set[nums[i]]=value
//		}
//		sum+=nums[i]
//	}
//
//	res:=[]int{sameNum,((1+length)*length/2)-(sum-sameNum)}
//	return res
//
//}

func findErrorNums(nums []int) []int {
	length:=len(nums)
	templist:=make([]int,length+1)

	res1,res2:=0,0

	for i:=0;i<length;i++ {
		templist[nums[i]]++
	}

	for i:=1;i<length+1;i++ {
		if templist[i] == 2 {
			res1 = i
		}else if templist[i] == 0 {
			res2 = i
		}
	}

	return []int{res1,res2}
}


