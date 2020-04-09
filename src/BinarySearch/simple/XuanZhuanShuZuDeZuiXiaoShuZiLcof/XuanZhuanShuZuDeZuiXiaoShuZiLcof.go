package XuanZhuanShuZuDeZuiXiaoShuZiLcof

func minArray(numbers []int) int {
    i:=0
    j:=len(numbers)-1
    for i<j {
        min:=i+(j-i)/2
        if numbers[min] > numbers[j] {
            i = min + 1
        }else if numbers[min] < numbers[j] {
            j = min
        }else {
            j --
        }
    }
    return numbers[i]
}