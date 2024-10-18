package v8

import (
	"testing"
	"testing/quick"
)

func TestPropertiesOfConversion(t *testing.T) {
	// 通过 quick.Check 测出了 负数会报错，过大的数会卡死，考虑实际的使用场景
	// 这里的 罗马数字本身就不能表示负数，所以原来使用的int 类型就有问题，罗马数字能表示的数据范围也有限，需要进一步限制
	// 如果实际的场景中不需要负数，则直接使用无符号类型 uint
	// 进一步限制数据范围 ，可以通过限制位数实现， uint16 最大为65535
	assertion := func(arabic uint16) bool {
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	// 默认运行 quick.Check 的次数是 100 ，可以修改配置更改
	if err := quick.Check(assertion, &quick.Config{MaxCount: 1000}); err != nil {
		t.Error("failed checks", err)
	}
}
