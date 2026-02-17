package util

import (
	"fmt"
)

func GenerateCode(prefix string, count int64) string {
	return fmt.Sprintf("%s%03d", prefix, count+1)
}
