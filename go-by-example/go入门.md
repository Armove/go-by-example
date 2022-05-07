go入门
go的函数
func add(a int ,b int) int {
    return a+b
}
go的函数返回多个值
func get(m map[string]int, k string) (v int, ok bool) {
    v, ok = m[k]
    return v, ok
}
go的指针
func add(n *int) {
    *n += 2
}

func main() {
    add(&n)
}
go的结构体
type user struct {
    name string
    password string
}
a := user(name:"zhang",password:"123456")
go的结构体方法
func (u user) checkpassword(password string) bool {
    return u.password == password
}

func main() {
    ok := a.checkpassword("123456")
}
go的错误处理
在返回值里加一个error
errors.New()
go的字符串操作
strings.Contains(,)
strings.Count()
strings.HasPrefix()
strings.HasSUffix()
等
字符串格式化
fmt.Println()
fmt.Printf("n=%v\n",n)
JSON处理
