package hello

import "testing"

func TestHello(t *testing.T) {

	// 在 Go 中，你可以在其他函数中声明函数并将它们分配给变量。你可以像调用普通函数一样调用它们。
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper() //t.Helper() 需要告诉测试套件这个方法是辅助函数（helper）。通过这样做，当测试失败时所报告的行号将在函数调用中而不是在辅助函数内部。
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	//t.Run("saying hello to people", func(t *testing.T) {
	//	got := Hello("Chris")
	//	want := "Hello, Chris"
	//	assertCorrectMessage(t, got, want)
	//})
	//
	//t.Run("empty string defaults to 'world'", func(t *testing.T) {
	//	got := Hello("")
	//	want := "Hello, World"
	//	assertCorrectMessage(t, got, want)
	//})

	// 现在需要支持第二个参数，指定问候的语言。如果一种不能识别的语言被传进来，就默认为英语。

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in English or unknown", func(t *testing.T) {
		got := Hello("Chris", "xxx")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Chris", "French")
		want := "Bonjour, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Chinese", func(t *testing.T) {
		got := Hello("张三", "Chinese")
		want := "你好， 张三"
		assertCorrectMessage(t, got, want)
	})
}
