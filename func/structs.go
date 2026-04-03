package _func

import (
	"encoding/json"
	"fmt"
)

// Account 首字母大写 public，小写 private
type Account struct {
	ID      int
	Name    string
	balance float64 // 小写字母开头，只能在当前的包访问
}

// NewAccount 构造函数 ，返回一个Account 类型的指针，避免大的结构体复制
func NewAccount(id int, name string) *Account {
	return &Account{
		ID:      id,
		Name:    name,
		balance: 0.0, // 初始化私有字段
	}
}

// GetBalance 获取账户余额 read-only operation, 无法修改原始的Account的值 照官方推荐：即使是只读方法，也使用指针接收者
func (a *Account) GetBalance() float64 {
	// a.balance = 9999.9 只会修改副本的值，不会修改原始的Account的值
	return a.balance
}

// SetBalance 设置账户余额 write operation, 可以修改原始的Account的值, 指针接受者
func (a *Account) SetBalance(b float64) {
	a.balance = b
}

// Deposit 存款 write operation, 可以修改原始的Account的值, 指针接受者
func (a *Account) Deposit(amount float64) {
	if amount > 0 {
		a.balance += amount
		fmt.Printf("存入 %.2f, 余额为 %.2f\n", amount, a.balance)
	}
}

type Logger struct {
	LogLevel string
}

func (l Logger) LogInfo(msg string) {
	fmt.Printf("[%s] %s\n", l.LogLevel, msg)
}

type Server struct {
	Logger    // 匿名嵌入 Logger 结构体，Server 可以直接调用 Logger 的方法
	Name      string
	IsRunning bool
}

func (s Server) Start() {
	s.IsRunning = true                    // 无法修改 Server 结构体的字段
	s.LogInfo(s.Name + "is starting ...") // 可以调用 Logger 的方法
}

func (s *Server) Stop() {
	s.IsRunning = false // 指针接受者，可以修改s的状态
	s.LogInfo(s.LogLevel + "stopped.")
}

type User struct {
	ID       int    `json:"user_id"`
	Username string `json:"username"`

	// 小写字段，不会被json.Marshal 序列化/反序列化
	password string `json:"password"`

	Metadata map[string]string `json:"metadata"`
}

func Structs() {
	account := NewAccount(1001, "张三")
	account.Deposit(1000)

	var acc2 Account
	fmt.Println("acc2:", acc2)

	acc3 := Account{
		ID:   1002,
		Name: "李四",
	}

	fmt.Println("acc3:", acc3)

	b := account.GetBalance()
	fmt.Println("当前余额:", b)

	account.SetBalance(10000)
	fmt.Println("设置余额后:", account.GetBalance())

	fmt.Println("account:", account)

	srv := Server{
		Logger: Logger{LogLevel: "INFO"},
		Name:   "WebApp Server",
	}

	fmt.Println("srv:", srv)

	fmt.Println("Logger's LogLevel:", srv.LogLevel)

	// 绝对路径
	fmt.Println("Logger's LogLevel:", srv.Logger.LogLevel)

	fmt.Println("Server's Name:", srv.Name)

	srv.Start()
	fmt.Printf("server status: %v\n", srv.IsRunning)
	srv.Stop()
	fmt.Printf("server status: %v\n", srv.IsRunning)

	// user
	u := User{
		ID:       101,
		Username: "master",
		password: "secret", // 私有字段
	}

	// 序列化 user 结构体为 JSON 字符串
	jsonData, _ := json.Marshal(u)
	//user json: {"user_id":101,"username":"master","metadata":null}
	fmt.Printf("user json: %s\n", string(jsonData))

	//var uZero User
	//uZero.Metadata["key"] = "value" // panic: assignment to entry in nil map

	uCorrect := User{
		ID:       102,
		Username: "safe_gopher",
		password: "123456", // 私有字段
		Metadata: make(map[string]string),
	}

	uCorrect.Metadata["email"] = "safe_gopher@example.com"
	uCorrect.Metadata["role"] = "admin"

	fmt.Println("struct: ", uCorrect)
}
