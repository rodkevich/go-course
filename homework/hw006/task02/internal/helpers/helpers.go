package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

// MultiplyOrUpperDepOfType
// 	Checks input if it’s an int, if true - returns input multiplied by 2
// 	If it’s not an integer, return uppercase input string
func MultiplyOrUpperDepOfType(items ...string) (rtn []string) {
	for i := 0; i < len(items); i++ {
		if x, err := strconv.Atoi(items[i]); err == nil {
			fmt.Printf("%q was treated as a number\n", items[i])
			rtn = append(rtn, strconv.Itoa(x*2))
			continue
		}
		rtn = append(rtn, strings.ToUpper(items[i]))
	}
	return
}
