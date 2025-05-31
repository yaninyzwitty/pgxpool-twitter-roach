package utils

import (
	"strconv"
	"strings"
)

func BuildServerAddr(port int) string {
	var b strings.Builder
	b.WriteByte(':')                  // Add colon
	b.WriteString(strconv.Itoa(port)) // Convert int to string
	return b.String()
}
