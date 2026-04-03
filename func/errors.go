package _func

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds 预定义哨兵错误, 不可变的错误值
var ErrInsufficientFunds = errors.New("insufficient funds in account")

func SimulateBankWithdrawal(balance, amount int) error {
	if amount > balance {
		return ErrInsufficientFunds
	}
	return nil
}

// Domo1 最基本的错误处理模式，使用预定义的 哨兵错误（Sentinel Error）来检查错误类型
func Domo1() {
	fmt.Println("哨兵错误：", ErrInsufficientFunds)

	// 模拟取款操作
	balance := 1000
	amount := 1500
	err := SimulateBankWithdrawal(balance, amount)
	if err != nil {
		if errors.Is(err, ErrInsufficientFunds) {
			fmt.Println("余额不足")
		} else {
			fmt.Println("其他错误:", err)
		}
	} else {
		fmt.Println("取款成功")
	}
}

var ErrDBConnection = errors.New("database connection failed")

func lowLevelDBCall() error {
	return ErrDBConnection
}

func businessLogic(userID int) error {
	err := lowLevelDBCall()
	if err != nil {
		return fmt.Errorf("user %d: cannot retrieve data from DB: %w", userID, err)
	}
	return nil
}

// Demo2 错误包装与解包 (%w 和 errors.Is) 当您需要添加上下文信息，但仍希望保留原始错误以便检查时，使用错误包装。
func Demo2() {
	err := businessLogic(404)

	if err != nil {
		fmt.Printf("业务逻辑错误: %v\n", err)
		if errors.Is(err, ErrDBConnection) {
			fmt.Println("数据库连接失败")
		} else {
			fmt.Println("其他业务逻辑错误")
		}
	} else {
		fmt.Println("业务逻辑执行成功")
	}
}

//自定义错误类型与提取信息 (errors.As)
//当错误需要携带额外数据（如 HTTP 状态码、请求ID）时，使用自定义结构体错误。

type APIError struct {
	HTTPStatus int
	RequestId  string
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("APIError: HTTP %d, RequestID: %s, Message: %s", e.HTTPStatus, e.RequestId, e.Message)
}

func getResource(id string) error {
	return &APIError{
		HTTPStatus: 404,
		RequestId:  "12345",
		Message:    fmt.Sprintf("resource %s not found", id),
	}
}

func Demo3() {
	wrappedErr := getResource("123")

	// errors.As 检查：检查错误链中是否包含 *APIError 类型，并将其赋值给 target 变量
	var apiErr *APIError // target 必须是指针类型

	if errors.As(wrappedErr, &apiErr) {
		fmt.Printf("API错误: HTTP %d, RequestID: %s, Message: %s\n", apiErr.HTTPStatus, apiErr.RequestId, apiErr.Message)
	} else {
		fmt.Println("其他错误")
	}
}

func Errors() {
	Domo1()
	Demo2()
	Demo3()
}
