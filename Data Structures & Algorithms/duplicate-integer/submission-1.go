func hasDuplicate(nums []int) bool {
    m := make(map[int]int)
    item := 0
    for i := 0; i < len(nums); i++ {
        item = nums[i]
        if _, ok := m[item]; ok {
            return true
        }
        m[item] = 1
    }
    return false
}
