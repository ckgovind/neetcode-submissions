func longestConsecutive(nums []int) int {
    ans := 0
    lMap := make(map[int]int)
    rMap := make(map[int]int)
    seen := make(map[int]bool)
    for _, num := range nums {
        if !seen[num] {
            seen[num] = true
            l := rMap[num-1]
            r := lMap[num+1]
            sum := l + r + 1
            lMap[num-l] = sum
            rMap[num+r] = sum
            if sum > ans {
                ans = sum
            }
        }
    }
    return ans
}