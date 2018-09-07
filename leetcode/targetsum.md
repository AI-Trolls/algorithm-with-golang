# target sum

- dfs 방식으로 접근
	- base case 검사 조건식에서 idx == len(nums)로 하는 바람에
	  같은 종료 케이스에 대해 2번씩 검사함.
	- pass 되긴 함, go라서 pass된듯, 2^20승의 로직이니 이렇게 풀면 안됨
```go
import "fmt"
func findTargetSumWays(nums []int, S int) int {
    var targetSum func (int, int, int) int 
    targetSum = func (idx int, ops int, prevTotal int) int {
        var total int
        switch ops {
        case 0:
            total = prevTotal - nums[idx]
        case 1:
            total = prevTotal + nums[idx]
        }
        
        if idx == len(nums) - 1 {
            if total == S {
                return 1
            } else {
                return 0
            }
        }
        
        return targetSum(idx + 1, 0, total) + targetSum(idx + 1, 1, total)
    }
    
    return targetSum(0, 0, 0) + targetSum(0, 1, 0)
}
```

- dfs + memo방식으로 접근
