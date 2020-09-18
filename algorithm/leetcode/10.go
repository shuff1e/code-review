package main

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	if len(s) == 0 {
		return true
	}

	firstMatch := (s[0] == p[0] || p[0] == '.')

	if len(p) >=2 && p[1]  == '*' {
		return isMatch(s,p[2:]) || (firstMatch && isMatch(s[1:],p))
	} else {
		return firstMatch && isMatch(s[1:],p[1:])
	}
}
