package utils

import "net/url"

func Escape(s string) string {
	// return strings.ReplaceAll(s, "\"", "\\\"")
	return url.QueryEscape(s)
}
