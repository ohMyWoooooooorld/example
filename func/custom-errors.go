package _func

import (
	"errors"
	"fmt"
	"time"
)

// ConnectionError 1. 定义结构体，通常以 Error 结尾
type ConnectionError struct {
	Host    string
	Port    int
	Timeout time.Duration
}

// 2. 必须实现 error 接口
func (e *ConnectionError) Error() string {
	return fmt.Sprintf("connection to %s:%d faild after %s", e.Host, e.Port, e.Timeout)
}

// 数中实例化自定义错误，并以 error 接口类型返回。
func connect(host string, port int) error {
	return &ConnectionError{ //必须返回指针，才能正确实现errors.As
		Host:    host,
		Port:    port,
		Timeout: 5 * time.Second,
	}
}

// IsTemporary 假设连接错误总是临时的
func (e *ConnectionError) IsTemporary() bool {
	return true
}

func CustomErrors() {
	var connErr *ConnectionError

	err := connect("localhost", 8080)

	// 检查 err 链中是否包含有 *ConnectionError 类型，如果有，赋值给 connErr
	if errors.As(err, &connErr) {
		fmt.Printf("Connection failed to port %d (Timeout: %s)", connErr.Port, connErr.Timeout)
		if connErr.IsTemporary() {
			// 自动重试
			fmt.Println("Retrying...")
		}
	} else if err != nil {
		fmt.Printf("Unknown error: %s", err.Error())
	}
}
