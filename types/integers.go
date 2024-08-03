package types

import "fmt"

func CheckInteger(integer int) string {
	msg := fmt.Sprintf("INTEGER %v", integer)
	return msg
}
