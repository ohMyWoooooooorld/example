package _func

import (
	"fmt"
)

// interface 是go实现多态和解耦合设计的关键
// 接口变量有两个信息：type，value
// 空接口 interface{} 是一个特殊的接口类型，它没有任何方法。 所有的类型都实现了空接口，因此可以将任何类型的变量赋值给空接口变量。

type Speaker interface {
	Speak() string
	Volume() int
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " is speaking"
}

func (d Dog) Volume() int {
	return 100
}

type Notifier interface {
	Send(message string)
}

type SMSNotifier struct {
	provider string
}

func (s SMSNotifier) Send(message string) {
	fmt.Printf("Sending SMS via %s: %s\n", s.provider, message)
}

type EmailNotifier struct {
	clientVersion string
}

func (e EmailNotifier) Send(message string) {
	fmt.Printf("Sending Email via %s: %s\n", e.clientVersion, message)
}

func SendUserAlert(n Notifier, user string, message string) {
	fullMsg := fmt.Sprintf("Alert for %s: %s", user, message)
	n.Send(fullMsg)
}

// any 是一个空接口类型，它可以表示任何类型的值
func printAnything(i any) {
	fmt.Printf("value: %v, Type: %T\n", i, i)
}

func Interfaces() {
	fmt.Println("interface design")

	smsClient := SMSNotifier{
		provider: "Twilio",
	}

	emailClient := EmailNotifier{
		clientVersion: "v2.1",
	}

	SendUserAlert(smsClient, "Alice", "Your account balance is low")
	SendUserAlert(emailClient, "Bob", "Your account balance is low")

	fmt.Println("接口实现多态特性")

	//空接口可以存储任何的值
	var a any = 42
	printAnything(a)

	var b any = "Go!!!"
	printAnything(b)

	printAnything(smsClient) //value: {Twilio}, Type: _func.SMSNotifier

	var d any = []int{1, 2, 3}
	printAnything(d)

	var e any = nil
	printAnything(e)

	// 正确：使用类型断言提取原始类型
	if s, ok := b.(string); ok {
		// 类型断言成功，s 是一个字符串
		fmt.Println("b is a string:", s)
	}
}
