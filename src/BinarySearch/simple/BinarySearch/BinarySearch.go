package BinarySearch

func search(nums []int, target int) int {
    left :=0
    right := len(nums)
    min :=0
    for left<right {
        min = left + (right -left)/2
        if nums[min] > target {
            right = min
        }else if nums[min] < target{
            left = min + 1
        }else {
            return min
        }
    }
    return -1
}