package services

import "strings"

var domainMap = map[string]string{
	"instagram.com": "ddinstagram.com",
	"pixiv.com":     "phixiv.net",
	"x.com":         "fixupx.com",
	"tiktok.com":    "vxtiktok.com",
	"twitter.com":   "fxtwitter.com",
	"reddit.com":    "rxddit.com",
	"imgur.com":     "s.imgur.com",
}

func AutoEmbed(urlStr string) string {
	urlStr = strings.Replace(urlStr, "www.", "", 1)
	for domain, replacement := range domainMap {
		urlStr = strings.Replace(urlStr, domain, replacement, 1)
	}
	return urlStr
}
