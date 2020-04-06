package PowerfulIntegers


func powerfulIntegers(x int, y int, bound int) []int {
	type void struct{}
	mp := make(map[int]void)

	var v int
	for i:=0;i<20 && Pow(x,i) <=bound;i++ {
		for j:=0;j<20 && Pow(y,j)<=bound;j++ {
			v = Pow(x,i)+Pow(y,j)
			if v <=bound {
				mp[v] = void{}
			}
		}
	}

	var res []int
	for key,_ := range mp {
		res = append(res,key)
	}

	return res
}

func Pow(x int,y int) int {
	sum:=1
	for i:=0;i<y;i++ {
		sum*=x
	}
	return sum
}