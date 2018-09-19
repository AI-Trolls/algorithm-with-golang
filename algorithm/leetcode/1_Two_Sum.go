// LeetCode Easy
// 1 Two Sum in Go Language
// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// Example:

// Given nums = [2, 7, 11, 15], target = 9,

// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

// Brute Force
// Time O(n^2), Space O(1)
func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i+1; j < len(nums); j++ {
            if (target - nums[i]) == nums[j] {
                return []int{i, j}
            } 
        }
	}
    return  []int{-1, -1}
}

// Two-pass Hash Table
// Time O(n), Space O(n)
func twoSum(nums []int, target int) []int {
    var m map[int]int
    m = make(map[int]int)
    for i := 0; i < len(nums); i++ {
        m[nums[i]] = i
    }
    for i := 0; i < len(nums); i++ {
        complement := target - nums[i]
        val, exists := m[complement]
        if exists == true && val != i {
            j := m[complement]
            return []int{i, j}
        }
    }
    return []int{-1, -1}
}

// One-pass Hash Table
// Time O(n), Space O(n)
func twoSum(nums []int, target int) []int {
    var m map[int]int
    m = make(map[int]int)
    for i := 0; i < len(nums); i++ {
        complement := target - nums[i]
        val, exists := m[complement]
        if exists == true && val != i {
            j := m[complement]
            return []int{i, j}
        }
        m[nums[i]] = i
    }
    return []int{-1, -1}
}
