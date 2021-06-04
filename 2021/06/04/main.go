package main

func openLock(deadends []string, target string) int {

	if target == "0000" {
		return 0
	}
	dead := map[string]int{}

	for _, v := range deadends {
		dead[v] = 0
	}

	if _, ok := dead["0000"]; ok {
		return -1
	}

	q := []string{"0000"}
	dead["0000"] = 0

	for len(q) != 0 {
		s := q[0]
		q = q[1:]
		depth := dead[s]

		for i := 0; i < 4; i++ {
			for d := -1; d <= 1; d += 2 {
				x := (int(s[i]-'0')+d+10)%10 + '0'
				n := s[0:i] + string(x) + s[i+1:]

				if _, ok := dead[n]; ok {
					continue
				}

				dead[n] = depth + 1
				q = append(q, n)

				if n == target {
					return dead[n]
				}
			}
		}

	}
	return -1
}
