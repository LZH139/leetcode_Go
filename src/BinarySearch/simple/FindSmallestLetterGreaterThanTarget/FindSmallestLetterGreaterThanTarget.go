package FindSmallestLetterGreaterThanTarget

func nextGreatestLetter(letters []byte, target byte) byte {
    left := 0
    right := len(letters)
    min := 0
    for left < right {
        min = left + (right -left)/2
        if letters[min] <= target {
            left = min+1
        }else {
            right = min
        }
    }
    return letters[left%len(letters)]
}