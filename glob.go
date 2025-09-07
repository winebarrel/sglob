package sglob

func Match(pat string, str string) bool {
	patScan := newScanner(pat)
	strScan := newScanner(str)

	for patScan.HasNext() {
		if atLeast, ok := patScan.SkipWildcard(); ok {
			nc, ok := patScan.Peek()

			if !ok {
				return strScan.RestLen() >= atLeast
			}

			n, ok := strScan.SkipBefore(nc)

			if !ok || n < atLeast {
				return false
			}
		} else {
			pc, _ := patScan.Next()
			sc, ok := strScan.Next()

			if !ok {
				return false
			}

			if pc != '?' && pc != sc {
				return false
			}
		}
	}

	return !strScan.HasNext()
}
