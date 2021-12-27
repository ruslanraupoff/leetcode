package main

func calculate(s string) int {
	var x, y, z int
	var c, d byte
	c = '+'

	idx := 0
	nextN := func() int {
		n := 0
		for idx < len(s) && s[idx] == ' ' {
			idx++
		}
		for idx < len(s) && '0' <= s[idx] && s[idx] <= '9' {
			n = n*10 + int(s[idx]-'0')
			idx++
		}
		return n
	}

	nextOpt := func() byte {
		opt := byte('+')
		for idx < len(s) && s[idx] == ' ' {
			idx++
		}
		if idx < len(s) {
			opt = s[idx]
			idx++
		}
		return opt
	}

	y = nextN()
	for idx < len(s) {
		d = nextOpt()
		z = nextN()

		if d == '*' || d == '/' {
			y = operate(y, z, d)
		} else {
			x = operate(x, y, c)
			c = d
			y = z
		}
	}

	return operate(x, y, c)
}

func operate(a, b int, opt byte) int {
	switch opt {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	default:
		return a / b
	}
}
