package maps

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {
	//dictionary := map[string]string{"test": "this is just a test"}
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("know word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNotFound

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		//由于 nil 指针异常，你永远不应该初始化一个空的 map 变量
		//var m map[string]string
		//相反，你可以初始化空 map，或使用 make 关键字创建 map：
		//dictionary = map[string]string{}
		//dictionary = make(map[string]string)
		dictionary := Dictionary{}
		key := "test"
		definition := "this is just a test"

		err := dictionary.Add("test", "this is just a test")

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "test"
		definition := "this is just a test"
		dictionary := Dictionary{key: definition}
		err := dictionary.Add("test", "this is just a test")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, key, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		key := "test"
		definition := "this is just a test"
		dictionary := Dictionary{key: definition}
		newDefinition := "new definition"

		_ = dictionary.Update(key, newDefinition)
		assertDefinition(t, dictionary, key, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		key := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Update(key, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	dictionary := Dictionary{key: "test definition"}

	dictionary.Delete(key)

	_, err := dictionary.Search(key)
	assertError(t, err, ErrNotFound)
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s' given", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if !errors.Is(got, want) {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, key, value string) {
	t.Helper()

	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertString(t, got, value)
}
