package main

func sortRGB1(s []string) []string {
	b := 0
	for i:=0 ; i<len(s); i++ {
		if s[i] != "B" {
			s[b] = s[i]
			b++
		}
	}

	for ; b < len(s); b++ {
		s[b] = "B"
	}

	r := len(s) - 1
	for i:= len(s)- 1 ; i >= 0; i-- {
		if s[i] != "R" {
			s[r] = s[i]
			r--
		}
	}

	for ; r >= 0; r-- {
		s[r] = "R"
	}
	return s
}

func sortRGB2(s []string) []string {
	// Dutch National Flag problem
	i, j := 0, 0
	l := len(s) - 1

	for j <= l {
		if s[j] == "R" {
			s[j], s[i] = s[i], s[j]
			i++
			j--
		} else if s[j] == "B" {
			s[j], s[l] = s[l], s[j]
			l--
		} else {
				j++
		}
	}
return s
}