package sumutil

import (
	"strconv"
	"fmt"
)

// SumlogHaiso thực hiện phép cộng hai số nguyên lớn được biểu diễn dưới dạng chuỗi.
func SumlogHaiso(Number1, Number2 string, MaxLen1, MaxLen2, MaxLengthNum , returnNumber int,result *[]string) {
	if MaxLengthNum < 0 {
		if returnNumber != 0 {
			*result = append(*result, strconv.Itoa(returnNumber))
		}
		return
	}
	var digit1, digit2 int
	if MaxLen1 < 0 {
		digit1 = 0
	} else {
		digit1, _ = strconv.Atoi(string(Number1[MaxLen1])) // Chuyển ký tự sang số
		MaxLen1--
	}

	if MaxLen2 < 0 {
		digit2 = 0
	} else {
		digit2, _ = strconv.Atoi(string(Number2[MaxLen2])) // Chuyển ký tự sang số
		MaxLen2--
	}

	NumTem := digit1 + digit2 + returnNumber

	fmt.Printf("%d + %d + %d= %d  ", digit1, digit2, returnNumber, NumTem)
	if NumTem >= 10 {
		returnNum := strconv.Itoa(NumTem)[0]
		returnNumber, _ = strconv.Atoi(string(returnNum))

	} else {
		returnNumber = 0
	}
	fmt.Printf(" viet %d ", NumTem%10)
	fmt.Printf("nho %d \n", returnNumber)

	*result = append(*result, strconv.Itoa(NumTem%10))
	SumlogHaiso(Number1, Number2, MaxLen1, MaxLen2, MaxLengthNum-1,returnNumber, result)

}
func ReverseSlice(result *[]string) {
	n := len(*result)
	for i := 0; i < n/2; i++ {
		(*result)[i], (*result)[n-1-i] = (*result)[n-1-i], (*result)[i]
	}
}
func Display(result *[]string) (int,  error) {
	str := ""
	for _, digit := range *result {
		str += digit
	}
	fmt.Println("Tong cac so viet hai so la: ", str)
	return strconv.Atoi(str)
}