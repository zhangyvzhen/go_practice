package pointer

import (
	"errors"
	"fmt"
)

// Bitcoin
// 我们需要一个比特币类型，但是单独创建一个结构体太过头
// int 的表现来说已经很好了，但问题是它不具有描述性
// Go 允许从现有的类型创建新的类型
// 语法是 type MyName OriginalType
type Bitcoin int

// 自定义Sting 方法，这个接口是在 fmt 包中定义的
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	// 在 Go 中，如果一个符号（例如变量、类型、函数等）是以小写符号开头，那么它在 定义它的包之外 就是私有的。
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// InsufficientFundsError
// 在 Go 中，错误是值
// 对于测试来说并不关心值的内容，并且值有修改的可能
// 因此我们可以将其重构为一个变量，并为其提供一个单一的事实来源
var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return InsufficientFundsError
	}

	w.balance -= amount

	return nil
}
