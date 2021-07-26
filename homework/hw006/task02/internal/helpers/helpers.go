package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

func MultiplyOrUpperDepOfType(items []string) (rtn []string) {
	for i := 0; i < len(items); i++ {
		if _, errAtoi := strconv.Atoi(items[i]); errAtoi == nil {
			fmt.Printf("%q was treated as a number\n", items[i])
			var a int
			_, errConv2Int := fmt.Sscanf(items[i], "%d", &a)
			if errConv2Int != nil {
				return rtn
			}
			rtn = append(rtn, strconv.Itoa(a*2))
			continue
		}
		rtn = append(rtn, strings.ToUpper(items[i]))
	}
	return
}
