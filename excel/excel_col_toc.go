package excel

var ExcelChar = []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func ConvertNumToChar(num int) (string, error) {
	if num < 27 {
		return ExcelChar[num], nil
	}
	k := num % 26
	if k == 0 {
		k = 26
	}
	v := (num - k) / 26
	col, err := ConvertNumToChar(v)
	if err != nil {
		return "", err
	}
	cols := col + ExcelChar[k]
	return cols, nil
}
