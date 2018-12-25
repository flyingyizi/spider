package spider

import "regexp"

//spiderOption paraments feed the to NewCollector
type spiderOption func(*Collector)

// UserAgent sets the user agent used by the Collector.
func UserAgent(ua string) spiderOption {
	return func(c *Collector) {
		c.userAgent = ua
	}
}

// MaxDepth limits the recursion depth of visited URLs.
func MaxDepth(depth int) spiderOption {
	return func(c *Collector) {
		c.MaxDepth = depth
	}
}

// AllowedDomains sets the domain whitelist used by the Collector.
func AllowedDomains(domains ...string) spiderOption {
	return func(c *Collector) {
		c.AllowedDomains = domains
	}
}

// ParseHTTPErrorResponse allows parsing responses with HTTP errors
func ParseHTTPErrorResponse() spiderOption {
	return func(c *Collector) {
		c.ParseHTTPErrorResponse = true
	}
}

// DisallowedDomains sets the domain blacklist used by the Collector.
func DisallowedDomains(domains ...string) spiderOption {
	return func(c *Collector) {
		c.DisallowedDomains = domains
	}
}

// DisallowedURLFilters sets the list of regular expressions which restricts
// visiting URLs. If any of the rules matches to a URL the request will be stopped.
func DisallowedURLFilters(filters ...*regexp.Regexp) spiderOption {
	return func(c *Collector) {
		c.DisallowedURLFilters = filters
	}
}

// URLFilters sets the list of regular expressions which restricts
// visiting URLs. If any of the rules matches to a URL the request won't be stopped.
func URLFilters(filters ...*regexp.Regexp) spiderOption {
	return func(c *Collector) {
		c.URLFilters = filters
	}
}

// AllowURLRevisit instructs the Collector to allow multiple downloads of the same URL
func AllowURLRevisit() spiderOption {
	return func(c *Collector) {
		c.allowURLRevisit = true
	}
}

// MaxBodySize sets the limit of the retrieved response body in bytes.
func MaxBodySize(sizeInBytes int) spiderOption {
	return func(c *Collector) {
		c.MaxBodySize = sizeInBytes
	}
}

// CacheDir specifies the location where GET requests are cached as files.
func CacheDir(path string) spiderOption {
	return func(c *Collector) {
		c.CacheDir = path
	}
}

// IgnoreRobotsTxt instructs the Collector to ignore any restrictions
// set by the target host's robots.txt file.
func IgnoreRobotsTxt() spiderOption {
	return func(c *Collector) {
		c.IgnoreRobotsTxt = true
	}
}

// ID sets the unique identifier of the Collector.
func ID(id uint32) spiderOption {
	return func(c *Collector) {
		c.ID = id
	}
}

// Async turns on asynchronous network requests.
func Async(a bool) spiderOption {
	return func(c *Collector) {
		c.Async = a
	}
}

// DetectCharset enables character encoding detection for non-utf8 response bodies
// without explicit charset declaration. This feature uses https://github.com/saintfish/chardet
func DetectCharset() spiderOption {
	return func(c *Collector) {
		c.DetectCharset = true
	}
}

// // Debugger sets the debugger used by the Collector.
// func Debugger(d debug.Debugger) spiderOption {
// 	return func(c *Collector) {
// 		d.Init()
// 		c.debugger = d
// 	}
// }
