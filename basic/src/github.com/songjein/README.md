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

### return multiple items
```go
func mmm(a int, b int) (int, int) {
	return a + b, a - b
}

sum, diff = mmm(1, 5)
```

### ananymous func
```go
func() {
	// ~~~~
}()

func(s string) {
	// ~~~~ 
}("hiiii")

r := func(a int, b int) {

}(6, 5)
```

### closure ; can save the flow of the func
```go
func calc() func(x int) int {
	a, b := 1, 2

	return func(x int) int {
		return a*x + b	
	}
}

f := calc()
f(1)
```

### defer
```go
defer func() {
	// ~~~~
}()
```

### recover, panic
```go
func f() {
	defere func() {
		s := recover()
		// ~~~~
	}()

	panic("Error")
}
```

### ptr
```go
var nPtr *int
```

### struct
```go
type Person struct {
	name string
	age int
}

func main() {
	var p Person
	p.name = "jein"
	p.age = 22
}
```

### receiver 
- 구조체에 메서드를 연결할 수 있고, 리시버를 통해 멤버 변수에 접근 가능하다
```go
func (p *Person) print() { // call by ref
	fmt.Println(p.name, p.age)
}
func (p Person) print() { // call by val
	fmt.Println(p.name, p.age)
}
```

### embedding
- 상속 흉내 가능
```go
type Student struct {
	Person // is-a
	grade int
}
```

### interface
- 인터페이스가 명시하고 있는 메서드만 구현한다면 같은 타입으로 취급하고 싶을 때
```go
type A struct {}

func (a A) Print() {
	fmt.Println("hello")
}

type B struct {}

func (b B) Print() {
	fmt.Println("babo")
}

type Printer interface {
	Print()
}

func main() {
	var a A
	var p Printer = h
	p.Print() // hello
	
	var b B
	p = b
	p.Print() // babo
}
```

### empty interface & type assertion
- empty interface는 타입에 상관없이 값을 저장할 수 있고, 이를 활용할 때 타입 캐스팅이 필요
```go
var empty_interface interface{} = 1

var num int = empty_interface.(int)

fmt.Println(num)
```
