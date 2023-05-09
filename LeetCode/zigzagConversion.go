package main

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	var rows []string
	for i := 0; i < numRows; i++ {
		rows = append(rows, "")
	}
	for i := 0; i < len(s); i++ {
		row := i % (2*numRows - 2)
		if row >= numRows {
			row = 2*numRows - 2 - row
		}
		rows[row] += string(s[i])
	}
	var result string
	for i := 0; i < numRows; i++ {
		result += rows[i]
	}
	return result
}

func main() {
	println(convert("PAYPALISHIRING", 3))
}
