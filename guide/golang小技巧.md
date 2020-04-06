# SET

​		相比于java中的hashset和python中的字典，目前Go的标准库中是没有有set这种数据结构，但是因为map的键是唯一的，这恰好和set的特性一致，所以我们依旧可以使用map来实现简单的set，同时可以使用  `_, ok := m[key]` 的语法效率高。

​		一个比较简单的想法是这样，利用布尔型作为map的value，这个value用来占位可以忽略掉，但是这样会对内存造成一定浪费

```go
set := make(map[string]bool) // New empty set
set["Foo"] = true            // Add
for k := range set {         // Loop
    fmt.Println(k)
}
delete(set, "Foo")    // Delete
size := len(set)      // Size
exists := set["Foo"]  // Membership
```

​		而在go中空结构体不占用内存，那我们可以用空结构体作为map的值

优化后的代码，如下：

```go
type void struct{}
var member void

set := make(map[string]void) // New empty set
set["Foo"] = member          // Add
for k := range set {         // Loop
    fmt.Println(k)
}
delete(set, "Foo")      // Delete
size := len(set)        // Size
_, exists := set["Foo"] // Membership
```

