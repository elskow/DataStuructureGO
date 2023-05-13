package main

func fullJustify(words []string, maxWidth int) []string {
	var res []string
	var i, j int
	for i < len(words) {
		j = i + 1
		length := len(words[i])
		for j < len(words) && length+len(words[j])+1 <= maxWidth {
			length += len(words[j]) + 1
			j++
		}
		var line string
		if j == len(words) {
			for k := i; k < j; k++ {
				line += words[k]
				if k != j-1 {
					line += " "
				}
			}
			for len(line) < maxWidth {
				line += " "
			}
		} else {
			space := maxWidth - length + j - i - 1
			if j-i == 1 {
				line = words[i]
				for k := 0; k < space; k++ {
					line += " "
				}
			} else {
				average := space / (j - i - 1)
				remain := space % (j - i - 1)
				for k := i; k < j; k++ {
					line += words[k]
					if k != j-1 {
						for l := 0; l < average; l++ {
							line += " "
						}
						if remain > 0 {
							line += " "
							remain--
						}
					}
				}
			}
		}
		res = append(res, line)
		i = j
	}
	return res
}

func main() {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	maxWidth := 16
	for _, v := range fullJustify(words, maxWidth) {
		println(v)
	}
}
