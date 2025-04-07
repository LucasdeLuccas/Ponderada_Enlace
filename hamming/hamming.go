package hamming

func IsValidBinary(s string) bool {
	for _, c := range s {
		if c != '0' && c != '1' {
			return false
		}
	}
	return true
}

func CalculateParityBits(m int) int {
	r := 0
	for (1 << r) < (m + r + 1) {
		r++
	}
	return r
}

func IsPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

func GenerateHammingCode(data string) string {
	m := len(data)
	r := CalculateParityBits(m)
	n := m + r
	code := make([]rune, n)
	dataIndex := 0

	for i := 1; i <= n; i++ {
		if IsPowerOfTwo(i) {
			code[i-1] = '0'
		} else {
			code[i-1] = rune(data[dataIndex])
			dataIndex++
		}
	}

	for i := 0; i < r; i++ {
		parityPos := 1 << i
		parity := 0
		for j := 1; j <= n; j++ {
			if j&parityPos != 0 {
				if code[j-1] == '1' {
					parity ^= 1
				}
			}
		}
		if parity == 1 {
			code[parityPos-1] = '1'
		}
	}
	return string(code)
}

func CorrectHammingCode(code string) string {
	n := len(code)
	codeRunes := []rune(code)
	r := 0
	for (1 << r) <= n {
		r++
	}
	
	syndrome := 0

	for i := 0; i < r; i++ {
		parityPos := 1 << i
		parity := 0
		for j := 1; j <= n; j++ {
			if j&parityPos != 0 {
				if codeRunes[j-1] == '1' {
					parity ^=1
				}
			}
		}
		if parity != 0 {
			syndrome += parityPos
		}
	}

	if syndrome != 0 && syndrome <= n {
		idx := syndrome - 1
		if codeRunes[idx] == '0' {
			codeRunes[idx] = '1'
		} else {
			codeRunes[idx] = '0'
		}
	}
	return string(codeRunes)
}

func ExtractData(code string) string {
	n := len(code)
	var data []rune
	for i := 1; i <= n; i++ {
		if !IsPowerOfTwo(i) {
			data = append(data, rune(code[i-1]))
		}
	}
	return string(data)
}
