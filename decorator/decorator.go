package decorator

import "strings"

// URLDecorator decorates url
type URLDecorator = func(string) string

//URLDecorated decorated
type URLDecorated = string

// URLDecorate decorated
func URLDecorate(c URLDecorated, ds ...URLDecorator) URLDecorated {
	decorated := c
	for _, decorator := range ds {
		decorated = decorator(decorated)
	}
	return decorated
}

// RemoveHTTP from a url
func RemoveHTTP(s string) string {
	return strings.Replace(strings.ToLower(s), "http://", "", 1)
}


// RemoveHTTPS from a url
func RemoveHTTPS(s string) string {
	return strings.Replace(strings.ToLower(s), "https://", "", 1)
}

// RemoveWWW from a url
func RemoveWWW(s string) string {
	return strings.Replace(strings.ToLower(s), "www.", "", 1)
}

