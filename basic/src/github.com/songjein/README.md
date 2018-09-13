# go cheat sheet


### import
```go
import _ "fmt"
import f "fmt" // alias
```

### if 
```go
if b, err := ioutil.ReadFile("./abc.txt"); err == nil {

}
```

### loop -> only 'for'
### break/continue can use LABEL
```go
LABEL1:
	for i, j := 0, 0; i < 10; i, j = i+1, j+3 {
		for k := 0; k < 100; k++ {
			if k == 5 {
				break LABEL1	
				//continue LABEL1	
			}	
		}
	}
```

### switch, fallthrough
```go
switch {
case i >= 1 && i < 5:
	~~~
	fallthrough
case i >= 100 && i < 200:
	~~~
}
```

### range
```go
a := [5]int{1, 2, 3, 4, 5}
for i, value := range a {
	// ~~~	
}
```

###	copy array
```go
a := [5]int{1, 2, 3, 4, 5}
b := a
```

### slice
```go
a := []int{1, 2, 3}
var b []int

b = a // reference
```

### copy slice
```go
a := []int{1, 2, 3, 4}
b := make([]int, 3)

copy(b, a)
```

### partial slice
```go
a := []int{1, 2, 3, 4, 5}
a[:3]
a[2:]
```

### map
```go
m := make(map[string]int)
m["aaa"] = 1
m["bbb"] = 2
delete(m, "aaa")
```

### map range
```go
m := map[string]int {
	"abc": 10,
	"def": 20,
}
for key, value := range m {
	// ~~~
}
```