package util

import "net/http"

func ParseCookies(rawCookies string) []*http.Cookie {
	header := http.Header{}
	header.Add("Cookie", rawCookies)
	request := http.Request{Header: header}
	return request.Cookies()
}
