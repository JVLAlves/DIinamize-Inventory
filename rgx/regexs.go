package regexs

import (
	"regexp"
)

var (
	RegexHDandMemory   = regexp.MustCompile(`\s*(\d{1,3}[.,]?\d{0,3})`)
	RegexAssettagDigit = regexp.MustCompile(`\d{5}`)
	RegexMacOS         = regexp.MustCompile(`\s*(^\d{2}\.\d+)`)
	RegexpHostnameWin  = regexp.MustCompile(`[[:upper:]]{2,}[.\-|[:alnum:])]+[[:alnum:]]+[^\s]`)
	RegexCPU           = regexp.MustCompile(`([[:upper:]]+[a-z]+\(.\)[[:alpha:]]*[()]*.*[^\s])`)
	RegexHDWin         = regexp.MustCompile(`\d+[^\s]`)
	RegexSOWin         = regexp.MustCompile(`([[:alpha:]]*\s?Windows\s.*[^\s])`)
	//comentário
	RegexMemoriaWin   = regexp.MustCompile(`Memória física total:\s+(\d+[.,]\d+)\s(\w+[^\s])`)
	RegexProgramasWin = regexp.MustCompile(`[^\s][A-z]+\s[[:graph:]]+\s?[[:alnum:]]*.*[^\s]`)
	RegexHd = regexp.MustCompile(`\d+.\d+`)
)
