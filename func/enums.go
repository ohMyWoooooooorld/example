package _func

import "fmt"

// ServerState 服务器状态
type ServerState int

// ServerState 服务器状态的枚举值
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// stateName 服务器状态的字符串表示
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

// ss ServerState 的字符串表示, String 方法实现了 fmt.Stringer 接口
func (ss ServerState) String() string {
	return stateName[ss]
}

// transition 转换到指定状态
func transition(s ServerState) ServerState {
	// 状态转换逻辑
	switch s {
	case StateIdle:

		// 从空闲状态转换到已连接状态
		return StateConnected
	case StateConnected, StateRetrying:
		// 从已连接状态或重试状态转换到空闲状态
		return StateIdle
	case StateError:
		// 从错误状态转换到重试状态
		return StateRetrying
	default:
		// 未知状态
		panic(fmt.Errorf("未知状态: %d", s))
	}
}

func Enums() {
	// ns是一个ServerState类型的变量
	ns := transition(StateIdle)
	fmt.Println("ns = ", ns)

	ns2 := transition(ns)
	fmt.Println("ns2 = ", ns2)
}
