func twoSum(nums []int, target int) []int {
    n := len(nums)
    if n == 2{
        return []int{0, 1}
    } else{
        for i, num1 := range nums {
            for j, num2 := range nums[i+1:]{
                if (num1 + num2 == target) {
                    return []int {i, (j+i+1)}
                }
            }
        }
    }
    return []int {0, 1}
}