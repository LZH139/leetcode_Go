package FirstBadVersion

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool{
    return false
}

func firstBadVersion(n int) int {
    left := 1
    right := n
    min := 0
    for left<right {
        min = left+(right-left)/2
        if isBadVersion(min) {
            right = min
        }else {
            left = min+1
        }
    }
    return left
}