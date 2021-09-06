package main

func slowestKey(releaseTimes []int, keysPressed string) byte {
	i, r, mx := 0, byte(0), -1
	for i < len(keysPressed) {
		k := keysPressed[i]
		if i == 0 {
			mx = releaseTimes[i]
			r = k
		} else {
			d := releaseTimes[i] - releaseTimes[i-1]
			if d > mx {
				mx = d
				r = k
			} else if d == mx && k > r {
				r = k
			}
		}
		i += 1
	}
	return r
}
