# gostlings Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 生成 gostlings —— 51 道 Go 可执行练习题（exercises/ 故意待修复）+ 51 份参考答案（solutions/ 全部可运行）+ README + 验证脚本。

**Architecture:** 单 Go module（`module gostlings`），每题独立目录 `exercises/<主题>/<题名>/main.go`（package main），solutions/ 同构。验证 = `go run`（14_testing 用 `go test`）。只用标准库。

**Tech Stack:** Go 1.23.2，bash。

**Spec:** `docs/superpowers/specs/2026-07-23-gostlings-design.md`（题目格式约定、四段式题头、题型定义以此为准）

**通用约定（每个 Task 都必须遵守）：**
- 题头四段式注释（英文）：`// Concept:` `// Task:` `// Expected output:` `// Hint:`（Hint 注明对应 Go Tour 章节），待改处用 `// TODO:` 标注
- **代码文件（含注释、字符串字面量、预期输出）全部用英文**；只有 README.md 用中文
- exercises/ 的题必须处于"未完成"状态：编译错误、panic、死锁或输出错误，均为预期；学生修复后 `go run` 输出须等于 "Expected output"
- solutions/ 同路径文件 = 修复后的完整可运行版本，题头注释保留
- 所有文件必须 `gofmt` 干净
- 注意：Go 1.22+ 循环变量每次迭代都是新变量，goroutines3 不得依赖旧语义出题

---

### Task 1: 脚手架（go.mod + check.sh）

**Files:**
- Create: `go.mod`
- Create: `check.sh`

- [ ] **Step 1: 写 go.mod**

```
module gostlings

go 1.23
```

- [ ] **Step 2: 写 check.sh**

```bash
#!/bin/sh
# Run every solution; list failures and exit non-zero if any fail
fail=0
for dir in solutions/*/*/; do
  d="./${dir%/}"
  case "$dir" in
    solutions/14_testing/*) cmd="go test" ;;
    *) cmd="go run" ;;
  esac
  if ! $cmd "$d" >/dev/null 2>&1; then
    echo "FAIL: $dir"
    fail=1
  fi
done
if [ "$fail" -eq 0 ]; then
  echo "All solutions pass ✓"
else
  exit 1
fi
```

- [ ] **Step 3: 验证**

Run: `cd /Users/wangxinalex/SelfStudy/Rust/gostlings && go mod verify && sh check.sh`
Expected: check.sh 输出 "All solutions pass ✓"（solutions 为空时循环不执行，直接通过）

- [ ] **Step 4: Commit**

```bash
git add go.mod check.sh
git commit -m "feat: 脚手架 go.mod + check.sh"
```

---

### Task 2: 00_intro（格式范例，后续 Task 照此风格）

**Files:**
- Create: `exercises/00_intro/intro1/main.go`
- Create: `exercises/00_intro/intro2/main.go`
- Create: `solutions/00_intro/intro1/main.go`
- Create: `solutions/00_intro/intro2/main.go`

- [ ] **Step 1: 写 4 个文件**

`exercises/00_intro/intro1/main.go`：

```go
// Concept: program structure and the fmt package
// Task: fix this program so it compiles and runs
// Expected output: Hello, gostlings!
// Hint: printing text requires importing the fmt package (Go Tour: Basics 1)

package main

func main() {
	// TODO: Fix the line below; it does not compile.
	Println("Hello, gostlings!")
}
```

`solutions/00_intro/intro1/main.go`：

```go
// Concept: program structure and the fmt package
// Task: fix this program so it compiles and runs
// Expected output: Hello, gostlings!
// Hint: printing text requires importing the fmt package (Go Tour: Basics 1)

package main

import "fmt"

func main() {
	fmt.Println("Hello, gostlings!")
}
```

`exercises/00_intro/intro2/main.go`：

```go
// Concept: package declaration and the main function
// Task: an executable program needs `package main` and `func main()` — add them
// Expected output: I can run!
// Hint: a Go program starts in the main function of package main (Go Tour: Basics 1)

// TODO: Add the package declaration on this line.

import "fmt"

// TODO: Add the main function that calls fmt.Println("I can run!")
```

`solutions/00_intro/intro2/main.go`：

```go
// Concept: package declaration and the main function
// Task: an executable program needs `package main` and `func main()` — add them
// Expected output: I can run!
// Hint: a Go program starts in the main function of package main (Go Tour: Basics 1)

package main

import "fmt"

func main() {
	fmt.Println("I can run!")
}
```

- [ ] **Step 2: 验证 solutions 通过**

Run: `go run ./solutions/00_intro/intro1 && go run ./solutions/00_intro/intro2`
Expected: 依次输出 `Hello, gostlings!` 和 `I can run!`

- [ ] **Step 3: 验证 exercises 按预期失败**

Run: `go run ./exercises/00_intro/intro1; go run ./exercises/00_intro/intro2`
Expected: 均编译失败（undefined: Println / expected 'package'）

- [ ] **Step 4: Commit**

```bash
git add exercises solutions
git commit -m "feat: 00_intro exercises and solutions"
```

---

### Task 3: 01_variables

**Files:** `exercises/01_variables/variables{1..4}/main.go` + `solutions/` 同构 4 个

题目规格（题型 / 破在哪 / Expected output / solution 关键行）：

- [ ] **Step 1: variables1**（修错）`x = 5` assigned without declaration (undefined: x)；TODO asks to declare x → prints `x has the value 5`。solution: replace bad line with `var x = 5`（或 `x := 5`）。
- [ ] **Step 2: variables2**（修错）`message := "first"` followed by `message := "second"` (redeclared)；TODO fix → prints `second`。solution: second line becomes `message = "second"`。
- [ ] **Step 3: variables3**（填空）main only has `fmt.Println("the zero value of count is", count)`；TODO declare int `count` without assigning → prints `the zero value of count is 0`。solution: `var count int`。
- [ ] **Step 4: variables4**（填空）const block has only `Red = iota`；TODO add Green、Blue so values are 1、2；main prints the three constants → `0 1 2`。solution: add bare `Green` and `Blue` lines。

- [ ] **Step 5: 验证**

Run: `for t in 1 2 3 4; do go run ./solutions/01_variables/variables$t; done` → outputs as above；4 个 exercises 均失败。

- [ ] **Step 6: Commit** `git add exercises/01_variables solutions/01_variables && git commit -m "feat: 01_variables exercises and solutions"`

---

### Task 4: 02_functions

**Files:** `functions{1..4}` × exercises+solutions

- [ ] **Step 1: functions1**（修错）main calls `sayHello()` which is undefined；TODO define it → prints `Hello from a function!`。solution: `func sayHello() { fmt.Println("Hello from a function!") }`。
- [ ] **Step 2: functions2**（修错）given `func add(a int, b int) int { return a + b }`，main calls `add(1, "2")` (type mismatch)；TODO fix the call → prints `3`。solution: `add(1, 2)`。
- [ ] **Step 3: functions3**（填空）given signature `func divide(a, b int) (int, int)`，body TODO returns quotient and remainder；main does `q, r := divide(7, 2)` → prints `3 1`。solution: `return a/b, a%b`。
- [ ] **Step 4: functions4**（填空）given signature `func sum(nums ...int) int`，body TODO；main prints `sum(1, 2, 3)` → `6`。solution: range-accumulate。

- [ ] **Step 5: 验证**（solutions 输出正确，exercises 失败）
- [ ] **Step 6: Commit** `git commit -m "feat: 02_functions exercises and solutions"`

---

### Task 5: 03_control_flow

**Files:** `if1, for1, switch1, defer1` × exercises+solutions

- [ ] **Step 1: if1**（修错）`n := 9`，`if n%2 == 0` prints "odd" else prints "even" — branches swapped；TODO fix so output is correct → prints `odd`。solution: change condition to `!= 0`（或交换分支）。
- [ ] **Step 2: for1**（填空）`sum := 0` and the print are given；TODO add a for loop adding 1..10 into sum → prints `55`。solution: `for i := 1; i <= 10; i++ { sum += i }`。
- [ ] **Step 3: switch1**（填空）`switch score := 50; {` with cases `>= 90` → "excellent"、`>= 60` → "pass"；TODO add the default branch → prints `fail`。solution: `default: fmt.Println("fail")`。
- [ ] **Step 4: defer1**（填空）header explains defers run LIFO；TODO use three defers so the program prints 3, 2, 1 (each on its own line)。solution: `defer fmt.Println(1)` then `defer fmt.Println(2)` then `defer fmt.Println(3)`。

- [ ] **Step 5: 验证**
- [ ] **Step 6: Commit** `git commit -m "feat: 03_control_flow exercises and solutions"`

---

### Task 6: 04_pointers

**Files:** `pointers{1..3}` × exercises+solutions

- [ ] **Step 1: pointers1**（修错）`n := 42; p := &n`，code prints `p` (an address)；TODO print the pointed-to value → prints `42`。solution: `fmt.Println(*p)`。
- [ ] **Step 2: pointers2**（填空）given `func setZero(p *int) { /* TODO */ }`，main does `n := 5; setZero(&n)` then prints n → `0`。solution: `*p = 0`。
- [ ] **Step 3: pointers3**（修错）`var p *int; fmt.Println(*p)` panics (nil dereference)；TODO fix with `new` → prints `0`。solution: `p := new(int)` then print `*p`。

- [ ] **Step 4: 验证**
- [ ] **Step 5: Commit** `git commit -m "feat: 04_pointers exercises and solutions"`

---

### Task 7: 05_slices

**Files:** `slices{1..4}` × exercises+solutions

- [ ] **Step 1: slices1**（填空）TODO create an int slice containing 1, 2, 3 assigned to `s`；`fmt.Println("length:", len(s))` → prints `length: 3`。solution: `s := []int{1, 2, 3}`。
- [ ] **Step 2: slices2**（修错）`s := []int{1, 2}` then `append(s, 3)` without using the result；prints `[1 2]`；TODO fix → prints `[1 2 3]`。solution: `s = append(s, 3)`。
- [ ] **Step 3: slices3**（填空）`s := []int{1, 2, 3, 4, 5}`；TODO use a slice expression to get `[2 3 4]` into `sub` and print it → `[2 3 4]`。solution: `sub := s[1:4]`。
- [ ] **Step 4: slices4**（修错）`a := []int{1, 2, 3}; b := a; b[0] = 100` then printing a gives `[100 2 3]`；TODO use copy to make b independent → prints `[1 2 3]`。solution: `b := make([]int, len(a)); copy(b, a)`。

- [ ] **Step 5: 验证**
- [ ] **Step 6: Commit** `git commit -m "feat: 05_slices exercises and solutions"`

---

### Task 8: 06_maps

**Files:** `maps{1..3}` × exercises+solutions

- [ ] **Step 1: maps1**（填空）`m := map[string]int{}` given；TODO add key `"apple"` with value 3；print `m["apple"]` → `3`。solution: `m["apple"] = 3`。
- [ ] **Step 2: maps2**（修错）`m := map[string]int{"a": 1}`，code does `v := m["b"]` then prints "b exists with value v" — wrong logic；TODO use the comma-ok idiom so a missing key prints `b does not exist`。solution: `if v, ok := m["b"]; ok { ... } else { fmt.Println("b does not exist") }`。
- [ ] **Step 3: maps3**（填空）`m := map[string]int{"a": 1, "b": 2}`；TODO delete `"a"` then range-print the rest → prints `b 2`。solution: `delete(m, "a")` + `for k, v := range m { fmt.Println(k, v) }`。

- [ ] **Step 4: 验证**
- [ ] **Step 5: Commit** `git commit -m "feat: 06_maps exercises and solutions"`

---

### Task 9: 07_structs

**Files:** `structs{1..3}` × exercises+solutions

- [ ] **Step 1: structs1**（填空）TODO define a `Person` struct (fields `Name string`、`Age int`)；main has `p := Person{Name: "Alice", Age: 18}` and prints it → `{Alice 18}`。solution: add the type definition。
- [ ] **Step 2: structs2**（修错）struct declared with field order `Age int; Name string`，but initialized as `Person{"Tom", 18}` (positional, type mismatch — compile error)；TODO switch to field-name initialization → prints `{18 Tom}`。solution: `Person{Name: "Tom", Age: 18}`。
- [ ] **Step 3: structs3**（填空）given `type Address struct { City string }` and `type Person struct { Name string; Addr Address }`；TODO initialize a Person and print the city → prints `Hangzhou`。solution: `p := Person{Name: "Bob", Addr: Address{City: "Hangzhou"}}; fmt.Println(p.Addr.City)`。

- [ ] **Step 4: 验证**
- [ ] **Step 5: Commit** `git commit -m "feat: 07_structs exercises and solutions"`

---

### Task 10: 08_methods

**Files:** `methods{1..3}` × exercises+solutions

- [ ] **Step 1: methods1**（填空）given `type Rectangle struct { W, H int }` and main with `r := Rectangle{2, 3}; fmt.Println(r.Area())`；TODO define the Area method → prints `6`。solution: `func (r Rectangle) Area() int { return r.W * r.H }`。
- [ ] **Step 2: methods2**（修错）`type Counter struct{ n int }` with `func (c Counter) Increment() { c.n++ }` — value receiver, increment lost；main prints 0 after calling；TODO fix → prints `1`。solution: pointer receiver `func (c *Counter) Increment()`。
- [ ] **Step 3: methods3**（填空）same directory defines its own Rectangle + Area；TODO define pointer-receiver method `Scale(f int)` multiplying W、H by f；main: `r := Rectangle{2, 3}; r.Scale(2)` then prints `r.Area()` → `24`。solution: `func (r *Rectangle) Scale(f int) { r.W *= f; r.H *= f }`。

- [ ] **Step 4: 验证**
- [ ] **Step 5: Commit** `git commit -m "feat: 08_methods exercises and solutions"`

---

### Task 11: 09_interfaces

**Files:** `interfaces{1..4}` × exercises+solutions

- [ ] **Step 1: interfaces1**（填空）given `type Speaker interface{ Speak() string }`、`type Dog struct{}`、`func announce(s Speaker) { fmt.Println(s.Speak()) }`，main calls `announce(Dog{})`；TODO make Dog implement Speaker → prints `Woof!`。solution: `func (d Dog) Speak() string { return "Woof!" }`。
- [ ] **Step 2: interfaces2**（填空）given `type Celsius float64`，main does `fmt.Println(Celsius(25))`；TODO implement `fmt.Stringer` → prints `25°C`。solution: `func (c Celsius) String() string { return fmt.Sprintf("%g°C", float64(c)) }`。
- [ ] **Step 3: interfaces3**（填空）`var v any = "hello"`；TODO assert v to string and print its length → prints `length: 5`。solution: `s := v.(string); fmt.Println("length:", len(s))`。
- [ ] **Step 4: interfaces4**（修错）`var v any = "hello"` with `n := v.(int)` — panics；TODO rewrite with comma-ok so a failed assertion prints `not an int`。solution: `if n, ok := v.(int); ok { fmt.Println(n) } else { fmt.Println("not an int") }`。

- [ ] **Step 5: 验证**
- [ ] **Step 6: Commit** `git commit -m "feat: 09_interfaces exercises and solutions"`

---

### Task 12: 10_errors

**Files:** `errors{1..4}` × exercises+solutions

- [ ] **Step 1: errors1**（填空）given `func divide(a, b int) (int, error)`，the b==0 branch is TODO；main calls with b=0 and checks err → prints `divisor must not be 0`。solution: `return 0, errors.New("divisor must not be 0")` + `if err != nil { fmt.Println(err); return }` in main。
- [ ] **Step 2: errors2**（修错）`func checkAge(age int) error` wrongly does `return nil` when age < 0；main passes -1, checks err, prints "age is valid" — wrong；TODO fix → prints `age must not be negative`。solution: `return errors.New("age must not be negative")`。
- [ ] **Step 3: errors3**（填空）given `var ErrNotFound = errors.New("record not found")`，query function TODO wraps it via `fmt.Errorf("query user: %w", ErrNotFound)`；main TODO checks with `errors.Is` → prints `user not found`。solution: `%w` wrap + `errors.Is(err, ErrNotFound)`。
- [ ] **Step 4: errors4**（填空）given `type NotFoundError struct{ ID int }` with `Error() string` (pointer receiver, matching solution code below)，function returns it；main TODO uses `errors.As` to extract the ID → prints `missing user ID: 42`。solution: `var nfe *NotFoundError; if errors.As(err, &nfe) { fmt.Println("missing user ID:", nfe.ID) }`。

- [ ] **Step 5: 验证**
- [ ] **Step 6: Commit** `git commit -m "feat: 10_errors exercises and solutions"`

---

### Task 13: 11_generics

**Files:** `generics{1..3}` × exercises+solutions

- [ ] **Step 1: generics1**（填空）TODO define generic `func Max[T cmp.Ordered](a, b T) T`；main prints `Max(3, 7)` → `7`。solution: import `cmp`；`if a > b { return a }; return b`。
- [ ] **Step 2: generics2**（填空）TODO define constraint `type Number interface { ~int | ~float64 }` and `func Sum[T Number](xs []T) T`；main prints `Sum([]float64{1.5, 2.5})` → `4`。solution: range-accumulate。
- [ ] **Step 3: generics3**（填空）TODO define generic type `Stack[T any]` (backed by `[]T`) with `Push(v T)` and `Pop() T`；main: int stack push 1、2 then pop → prints `2`。solution: slice append / take last。

- [ ] **Step 4: 验证**
- [ ] **Step 5: Commit** `git commit -m "feat: 11_generics exercises and solutions"`

---

### Task 14: 12_goroutines

**Files:** `goroutines{1..3}` × exercises+solutions

- [ ] **Step 1: goroutines1**（修错）`func say(s string) { fmt.Println(s) }`，main calls `say("hello")` directly；TODO turn it into a goroutine and make sure it actually prints → prints `hello`。solution: `go say("hello")` + `time.Sleep(100 * time.Millisecond)`。
- [ ] **Step 2: goroutines2**（填空）main starts 3 goroutines in a loop, `var wg sync.WaitGroup` given；TODO add Add/Done/Wait → prints 3 lines `worker N done` (N=0..2, any order)。solution: `wg.Add(1)` before each go、`defer wg.Done()` inside、`wg.Wait()` at the end。
- [ ] **Step 3: goroutines3**（填空）`for i := 0; i < 3; i++` loop；TODO pass i as an argument to the goroutine (lesson: passing values is the safe habit) → prints 0、1、2 (any order)。solution: `go func(n int) { fmt.Println(n) }(i)` plus a WaitGroup。

- [ ] **Step 4: 验证**
- [ ] **Step 5: Commit** `git commit -m "feat: 12_goroutines exercises and solutions"`

---

### Task 15: 13_channels

**Files:** `channels{1..4}` × exercises+solutions

- [ ] **Step 1: channels1**（修错）main does `ch := make(chan string)` then sends `ch <- "hi"` and receives in the same goroutine — deadlock；TODO fix → prints `hi`。solution: move the send into `go func() { ch <- "hi" }()`。
- [ ] **Step 2: channels2**（修错）`ch := make(chan int)` (unbuffered)，main sends 1、2 sequentially before receiving — deadlock；TODO fix by changing only the make line → prints `1` and `2` on separate lines。solution: `ch := make(chan int, 2)`。
- [ ] **Step 3: channels3**（修错）goroutine sends 1、2、3 into ch but never closes；main `for v := range ch` prints then deadlocks；TODO add close → prints `1 2 3` on separate lines。solution: `close(ch)` after sending。
- [ ] **Step 4: channels4**（填空）two `chan string`，c1 already has a value ready、c2 never fires；TODO use select to print the first one that arrives → prints `received: fast lane`。solution: `select { case v := <-c1: fmt.Println("received:", v); case v := <-c2: fmt.Println(v) }`。

- [ ] **Step 5: 验证**（exercises 里死锁题 `go run` 会报 `fatal error: all goroutines are asleep - deadlock!`，属预期失败）
- [ ] **Step 6: Commit** `git commit -m "feat: 13_channels exercises and solutions"`

---

### Task 16: 14_testing

**Files:** `testing{1..3}/` 下各 `被测文件.go` + `被测文件_test.go` × exercises+solutions（共 12 个文件）。exercises 的测试必须 `go test` 失败，solutions 必须 `go test` 通过。

- [ ] **Step 1: testing1**（修错）`add.go` provides `func Add(a, b int) int { return a + b }`；`add_test.go` has `TestAdd` with `if Add(2, 3) != 6 { t.Error(...) }` — wrong expected value；TODO fix it → `go test` passes。solution: expected value becomes 5。题头注释注明用 `go test` 而非 `go run` 验证。
- [ ] **Step 2: testing2**（填空）`multiply.go` provides `func Multiply(a, b int) int`；`multiply_test.go` has the table-driven skeleton (cases slice + range loop)，TODO fill in at least 3 cases（`2*3=6`、`0*5=0`、`-2*3=-6`）→ `go test` passes。solution: complete the cases。
- [ ] **Step 3: testing3**（填空）`even.go` provides `func IsEven(n int) bool`；`even_test.go` TODO write two subtests with `t.Run`（"even" asserts IsEven(4) is true，"odd" asserts IsEven(3) is false）→ `go test -v` shows both subtests pass。solution: two t.Run blocks。

- [ ] **Step 4: 验证**

Run: `for t in 1 2 3; do go test ./solutions/14_testing/testing$t; done` → all ok；`go test ./exercises/14_testing/testing1` → FAIL（预期）。

- [ ] **Step 5: Commit** `git commit -m "feat: 14_testing exercises and solutions"`

---

### Task 17: README + 全量验证

**Files:**
- Create: `README.md`

- [ ] **Step 1: 写 README.md**（README 用中文；题目本身是英文）

```markdown
# gostlings

仿 [rustlings](https://github.com/rust-lang/rustlings) 的 Go 入门练习题：每题一个小程序，
修好它、跑通它，你就学会了这个知识点。

## 前置要求

Go 1.23+（`go version` 确认）

## 怎么做题

1. 按下表顺序做题
2. 打开 `exercises/<主题>/<题名>/main.go`，读题头注释（Concept / Task / Expected output / Hint）
3. 修改 `// TODO:` 标记的地方
4. 运行验证：

   ```sh
   go run ./exercises/01_variables/variables1
   ```

   （14_testing 主题用 `go test ./exercises/14_testing/testing1`）
5. 输出符合预期 = 过关。卡住了看 `solutions/` 里的同路径答案

## 做题顺序与 Tour 章节映射

| 主题 | 题数 | Go Tour 章节 |
|---|---|---|
| 00_intro | 2 | Basics 1 |
| 01_variables | 4 | Basics 8-12 |
| 02_functions | 4 | Basics 4-7 |
| 03_control_flow | 4 | Flowcontrol 1-12 |
| 04_pointers | 3 | Moretypes 1-5 |
| 05_slices | 4 | Moretypes 7-12 |
| 06_maps | 3 | Moretypes 19-23 |
| 07_structs | 3 | Moretypes 2-5 |
| 08_methods | 3 | Methods 1-6 |
| 09_interfaces | 4 | Methods 9-15 |
| 10_errors | 4 | Methods 19-20 |
| 11_generics | 3 | Generics 1-2 |
| 12_goroutines | 3 | Concurrency 1 |
| 13_channels | 4 | Concurrency 2-5 |
| 14_testing | 3 | （Testing：go.dev/doc/tutorial/add-a-test） |

## 校验全部答案

```sh
sh check.sh   # 跑通 solutions/ 下全部 51 题
```
```

- [ ] **Step 2: 全量验证**

Run: `sh check.sh` → `All solutions pass ✓`；`gofmt -l .` → 无输出；逐主题数一遍核对 15 个主题、exercises 与 solutions 各 51 题齐全。

- [ ] **Step 3: Commit**

```bash
git add README.md
git commit -m "docs: README 做题指南"
```

---

## Self-Review 记录

- Spec 覆盖：51 题全部对应 Task 2-16；README/check.sh/验证对应 Task 1、17 ✓
- 占位符：题目内容以"题型/破在哪/预期输出/solution 关键行"四要素锁定，执行者可直接成文 ✓
- 一致性：题名与 spec 表一致；Task 10 methods3 已注明"Area 在本目录内自带定义"避免跨题依赖；Task 12 errors4 注明 Error() 为指针接收者以匹配 errors.As 用法 ✓
- 语言：代码文件全英文（用户 2026-07-23 追加要求），README 中文 ✓
