/*
 * Given an array of integers, return indices of the two numbers such that they add up to a specific target.
 *
 * You may assume that each input would have exactly one solution, and you may not use the same element twice.
 *
 * Example:
 * Given nums = [2, 7, 11, 15], target = 9,
 *
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 */

func twoSum(nums []int, target int) []int {
    var res []int
    for i1 := 0; i1 < len(nums); i1++ {
        exists, i2 := check(nums[i1+1:], target-nums[i1])
        if exists {
            return []int{i1, i2+i1+1}
            break
        }
    }
	//should never get here
    return res
}

//use a binary search tree to make it faster
func check(nums []int, otherTarget int) (bool, int) {
    for i := range nums {
        if nums[i] == otherTarget {
            return true, i
        }
    }
    return false, 0
}