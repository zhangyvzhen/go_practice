package pointer

import (
	"errors"
	"testing"
)

// 声明一个 Wallet（钱包）结构体，我们可以利用它存放 Bitcoin（比特币）
func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)

	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, Bitcoin(20))
		// 测试 报错
		assertError(t, err, InsufficientFundsError)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
func assertError(t *testing.T, got error, want error) {
	// nil 是其他编程语言的 null。错误可以是 nil，因为返回类型是 error，这是一个接口
	// 如果你看到一个函数，它接受参数或返回值的类型是接口，它们就可以是 nil
	if got == nil {
		//t.Fatal。如果它被调用，它将停止测试。
		//这是因为我们不希望对返回的错误进行更多断言。
		//如果没有这个，测试将继续进行下一步，并且因为一个空指针而引起 panic。
		t.Fatal("didn't get an error but wanted one")
	}

	if !errors.Is(got, want) {
		t.Errorf("got {%s}, want {%s}", got, want)
	}
}
