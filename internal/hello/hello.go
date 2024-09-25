package hello

const spanish = "Spanish"
const french = "French"
const chinese = "Chinese"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const chineseHelloPrefix = "你好， "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// 在我们的函数签名中，我们使用了 命名返回值（prefix string）。
//
// 这将在你的函数中创建一个名为 prefix 的变量。
//
// 它将被分配「零」值。这取决于类型，例如 int 是 0，对于字符串它是 ""。
//
// 你只需调用 return 而不是 return prefix 即可返回所设置的值。
//
// 这将显示在 Go Doc 中，所以它使你的代码更加清晰。
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case chinese:
		prefix = chineseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
