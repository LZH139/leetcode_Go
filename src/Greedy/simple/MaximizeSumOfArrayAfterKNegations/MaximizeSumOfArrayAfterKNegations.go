package MaximizeSumOfArrayAfterKNegations

func largestSumAfterKNegations(A []int, K int) int {
    var s = [201]int{}
    sum:=0
    for _,value := range A {
        s[value+100]++
    }
    temp := 0
    for key,value := range s {
        if value > 0 {
            if K > 0 && key < 100 {
                if K >= value {
                    sum = sum-(key-100) * value
                    K-=value
                }else {
                    sum = sum-(key-100) * K + (key-100) * (value- K)
                    K=0
                }

            }else if K > 0 && key >= 100 {
                if K & 1 == 0 {
                    sum+= (key-100) * value

                }else {
                    if -(temp-100) < key-100 {
                        sum = sum + (temp-100)*2 + (key-100)*value
                    }else {
                        sum = sum - (key-100) * 1 + (key-100) * (value-1)
                    }
                }
                K = 0
            }else {
                sum+=(key-100) * value
            }
            temp = key
        }

    }
    return sum

}