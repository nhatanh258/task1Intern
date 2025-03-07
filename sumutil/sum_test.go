package sumutil

import (
	"fmt"
	"strings"
	"testing"
)

func TestSumlogHaiso(t *testing.T) {
	tests := []struct {
		num1, num2 string
		expected   string
	}{
		{"123", "456", "579"},
		{"999", "1", "10004"},
		{"0", "0", "4"},
		{"56789", "43211", "100000"},
	}

	for _, test := range tests {
		var result []string
		SumlogHaiso(test.num1, test.num2, len(test.num1)-1, len(test.num2)-1, maxLen(test.num1, test.num2), 0, &result)
		ReverseSlice(&result)
		actual := strings.Join(result, "")// Chuyển []string thành string

		if actual != test.expected {
			t.Errorf("SumlogHaiso(%s, %s) = %s; khong bang %s", test.num1, test.num2, actual, test.expected)
		}else {
			fmt.Printf("Test thành công với số 1: %s, số 2: %s, kết quả: %s\n", test.num1, test.num2, actual)
		}
	}
}
func maxLen(n1, n2 string) int {
	if len(n1) > len(n2) {
		return len(n1) - 1
	}
	return len(n2) - 1
}
