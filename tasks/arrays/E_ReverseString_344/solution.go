package E_ReverseString_344

func reverseString(s []byte) []byte {

	for i := 0; i < len(s)/2; i++ {
		start := s[i]
		j := len(s) - 1 - i
		end := s[j]
		if start != end {
			start := s[i]
			end := s[j]
			s[i] = end
			s[j] = start
		}

	}
	return s
}
