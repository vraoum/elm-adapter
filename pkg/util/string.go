package util

func SplitSize(str string, splitNumber int) []string {
	res := ""
	a := []rune(str)
	var arr []string
	for i, r := range a {
		res = res + string(r)
		if i > 0 && (i+1)%splitNumber == 0 {
			arr = append(arr, res)
			res = ""
		}
	}
	if res != "" {
		arr = append(arr, res)
	}
	return arr
}
