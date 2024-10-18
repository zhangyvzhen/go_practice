package v2

func ConvertToRoman(num int) string {
	var s string
	switch num {
	case 1:
		s = "I"
	case 2:
		s = "II"
	}
	return s
}
