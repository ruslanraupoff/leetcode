package main

func threeEqualParts(arr []int) []int {
	n, s := len(arr), 0

	for i := 0; i < n; i++ {
		s += arr[i]
	}

	if s == 0 {
		return []int{0, n - 1}
	}

	if s%3 != 0 {
		return []int{-1, -1}
	}

	k, j, t, i := s/3, n-1, 0, 0

	for j >= 0 {
		if arr[j] == 1 {
			t += 1
		}
		if t == k {
			break
		}
		j -= 1
	}

	for i < n {
		if arr[i] == 1 {
			break
		}
		i += 1
	}

	for l := j; l < n; l++ {
		if arr[i] != arr[l] {
			return []int{-1, -1}
		}
		i += 1
	}

	k = i
	for k < n {
		if arr[k] == 1 {
			break
		}
		k += 1
	}

	for l := j; l < n; l++ {
		if arr[k] != arr[l] {
			return []int{-1, -1}
		}
		k += 1
	}

	return []int{i - 1, k}
}
