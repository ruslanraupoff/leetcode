package main

func addOperators(num string, target int) []string {
	res := []string{}

	var dfs func(string, string, int, int)

	dfs = func(s, t string, r, p int) {
		var cs, ns string
		var cn int

		if len(s) == 0 {
			if r == target {
				res = append(res, t)
			}
			return
		}

		for i := 1; i <= len(s); i++ {
			cs = s[:i]
			if cs[0] == '0' && len(cs) > 1 {
				return
			}

			cn = integer(cs)
			ns = s[i:]

			if len(t) == 0 {
				dfs(ns, cs, cn, cn)
			} else {
				dfs(ns, t+"+"+cs, r+cn, cn)
				dfs(ns, t+"-"+cs, r-cn, -cn)
				dfs(ns, t+"*"+cs, r-p+p*cn, p*cn)
			}
		}
	}

	dfs(num, "", 0, 0)

	return res
}

func integer(s string) int {
	res := int(s[0] - '0')
	for i := 1; i < len(s); i++ {
		res = res*10 + int(s[i]-'0')
	}
	return res
}
