package maps

import "errors"

// 将错误声明为常量，这需要我们创建自己的 DictionaryErr 类型来实现 error 接口
// 它使错误更具可重用性和不可变性
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	value, exists := d[key]
	if !exists {
		return "", ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	// ; 是在分割语句, 取exists 作为if的条件
	if _, exists := d[key]; exists {
		return ErrWordExists
	}
	d[key] = value
	return nil
}

func (d Dictionary) Update(key string, definition string) error {
	_, err := d.Search(key)

	switch {
	case errors.Is(err, ErrNotFound):
		//我们可以重用 ErrNotFound 而不添加新错误。但是，更新失败时有更精确的错误通常更好
		//特定的错误描述可以为你提供有关错误的更多信息。以下是一个 Web 应用中的示例：
		//遇到 ErrNotFound 时可以重定向用户，但遇到 ErrWordDoesNotExist 时会显示错误消息。
		return ErrWordDoesNotExist
	case err == nil:
		d[key] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
