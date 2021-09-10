package private_utils

import (
	"bytes"
	"fmt"
	"strings"
)

// [1, 2, 3] -> "1, 2, 3" sep = ,
func SliceToString(s []interface{}, sep string) string {
	return strings.Replace(strings.Trim(fmt.Sprint(s), "[]"), " ", sep, -1)
}

// [1, 2, 3] -> "(1, 2, 3)"  sep = ,
func SliceToString2(s []interface{}, sep string) string {
	b := new(bytes.Buffer)
	b.WriteString("(")
	b.WriteString(SliceToString(s, sep))
	b.WriteString(")")
	return b.String()
}
