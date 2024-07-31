func isNumber(s string) bool {
    str := `^[+-]?((\d+(\.\d*)?)|(\.\d+))([eE][+-]?\d+)?$`
	res := regexp.MustCompile(str)
	return res.MatchString(s)
}