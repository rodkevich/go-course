package helpers

// ProcessInputAsType ...
func ProcessInputAsType(inp interface{}) (i int, s string, err error) {
	switch inp.(type) {
	case string:
		var rtn string
		print("case string:")
		return 0, rtn, nil
	case int:
		var rtn int
		print("case int:")
		return rtn, "", nil
	}
	return 0, "", nil
}

