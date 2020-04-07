package GuessNumberHigherOrLower

/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guess(num int) int{
    return -1
}

func guessNumber(n int) int {
    left := 1
    right := n
    min:=0
    temp := 0
    for left<right {
        min = left + (right-left)/2
        temp = guess(min)
        if temp > 0 {
            left = min+1
        }else if temp < 0{
            right = min-1
        }else {
            return min
        }
    }
    return -1
}