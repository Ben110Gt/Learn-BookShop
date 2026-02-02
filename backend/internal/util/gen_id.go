package util

import "fmt"

func FormatUserID(id uint) string {
	return fmt.Sprintf("U%03d", id)
}
