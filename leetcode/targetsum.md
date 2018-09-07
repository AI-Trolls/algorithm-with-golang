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
	- memo [20][2][1001]int 로 선언하니, 세번째 idx(total)이 음수가 나오는경우 런타임에러
	- memo map[string]int 로 선언하고 세 개의 정수를 하나의 스트링으로 묶어 저장하니 1/2정도로 빨라짐
```go
import "fmt"
import "strconv"

func findTargetSumWays(nums []int, S int) int {
    
    memo := make(map[string]int)
    
    var targetSum func (int, int, int) int 
    targetSum = func (idx int, ops int, prevTotal int) int {
        var total int
        switch ops {
        case 0:
            total = prevTotal - nums[idx]
        case 1:
            total = prevTotal + nums[idx]
        }
        
        key := strconv.Itoa(idx) + "-" + strconv.Itoa(ops) + "-" + strconv.Itoa(total)
        if val, ok := memo[key]; ok {
            return val
        }
        
        if idx == len(nums) - 1 {
            if total == S {
                return 1
            } else {
                return 0
            }
        }
        
        key = strconv.Itoa(idx) + "-" + strconv.Itoa(ops) + "-" + strconv.Itoa(total)
        memo[key] = targetSum(idx + 1, 0, total) + targetSum(idx + 1, 1, total)
        return targetSum(idx + 1, 0, total) + targetSum(idx + 1, 1, total)
    }
    
    return targetSum(0, 0, 0) + targetSum(0, 1, 0)
}
```
- 다른 사람 소스중 굉장히 간결했던 소스
```go
func findTargetSumWays(nums []int, S int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        ways := 0
        if nums[0] == S {
            ways++
        }
        if nums[0] == -S {
            ways++
        }
        return ways
    }
    return findTargetSumWays(nums[1:], S-nums[0]) + findTargetSumWays(nums[1:], S+nums[0])
}
```

- dp로 접근해보기
