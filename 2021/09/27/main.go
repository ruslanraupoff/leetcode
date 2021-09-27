package main

import "strings"

func numUniqueEmails(emails []string) int {
	h := make(map[string]bool, 100)
	for _, email := range emails {
		h[clean(email)] = true
	}
	return len(h)
}

func clean(email string) string {
	i := strings.IndexByte(email, '@')
	user, atDomain := email[:i], email[i:]
	return deleteDot(trim(user)) + atDomain
}

func deleteDot(username string) string {
	return strings.Replace(username, ".", "", -1)
}

func trim(username string) string {
	i := strings.IndexByte(username, '+')
	if i == -1 {
		return username
	}
	return username[:i]
}
