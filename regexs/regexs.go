package regexs

import (
	"regexp"
)

var (
	RegexHDandMemory   = regexp.MustCompile(`\s*(\d{1,3}[.,]?\d{0,3})`)
	RegexAssettagDigit = regexp.MustCompile(`\d{5}`)
	RegexMacOS         = regexp.MustCompile(`\s*(^\d{2}\.\d+)`)
)
