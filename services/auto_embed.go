package services

import "strings"

func AutoEmbed(url string) string {
	domainMap := map[string]string{
		"instagram.com": "ddinstagram.com",
		"pixiv.com":     "phixiv.net",
		"x.com":         "fixupx.com",
		"tiktok.com":    "vxtiktok.com",
		"twitter.com":   "fxtwitter.com",
		"reddit.com":    "rxddit.com",
		"imgur.com":     "s.imgur.com",
	}

	url = strings.Replace(url, "www.", "", -1)

	for original, replacement := range domainMap {
		if strings.Contains(url, original) {
			return strings.Replace(url, original, replacement, 1)
		}
	}
	return url
}
