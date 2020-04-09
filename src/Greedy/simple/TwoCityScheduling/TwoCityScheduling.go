package TwoCityScheduling

import "sort"

//func twoCitySchedCost(costs [][]int) int {
//   // [1 , 5 ,3 ,4]
//   //贪心算法思想。先优先A的最优的方式。剩下的都是B,由于没有排序。自己撸个insertingSort
//   costs = insertingSort(costs)
//   ans := 0
//   n := len(costs) / 2
//   for i := 0; i < n; i++ {
//       ans += costs[i][0] + costs[i+n][1]
//   }
//   return ans
//}
////插入排序
//func insertingSort(costs [][]int) [][]int {
//   for i:=1; i < len(costs) ;i++ {
//       for j :=0; j < i; j++ {
//           //swap
//           if (costs[i][0] - costs[i][1]) < (costs[j][0] - costs[j][1]) {
//               costs[i], costs[j] = costs[j], costs[i]
//           }
//       }
//   }
//   return costs
//}

func twoCitySchedCost(costs [][]int) int {
    l := len(costs)
    c := make([]int, l)
    res := 0
    for i := 0; i < l; i ++ {
        c[i] = costs[i][0] - costs[i][1]
        res += costs[i][1]
    }
    sort.Ints(c)
    for i := 0; i < l/2; i++ {
        res += c[i]
    }
    return res
}