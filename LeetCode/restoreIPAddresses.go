package main

import (
	"strings"
)

func restoreIpAddresses(s string) []string {
	result := []string{}
	restoreIpAddressesHelper(s, 0, []string{}, &result)
	return result
}

func restoreIpAddressesHelper(s string, index int, current []string, result *[]string) {
	if index == len(s) && len(current) == 4 {
		*result = append(*result, strings.Join(current, "."))
		return
	}

	if len(current) == 4 {
		return
	}

	for i := index; i < len(s); i++ {
		if i > index && s[index] == '0' {
			break
		}

		if isValid(s[index : i+1]) {
			current = append(current, s[index:i+1])
			restoreIpAddressesHelper(s, i+1, current, result)
			current = current[:len(current)-1]
		}
	}
}

func isValid(s string) bool {
	if len(s) > 3 {
		return false
	}

	if len(s) > 1 && s[0] == '0' {
		return false
	}

	if len(s) == 3 && s[0] > '2' {
		return false
	}

	if len(s) == 3 && s[0] == '2' && s[1] > '5' {
		return false
	}

	if len(s) == 3 && s[0] == '2' && s[1] == '5' && s[2] > '5' {
		return false
	}

	return true
}

func main() {
	s := "25525511135"
	for _, ip := range restoreIpAddresses(s) {
		println(ip)
	}
}
