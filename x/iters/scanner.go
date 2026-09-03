package iters

import (
	"bufio"
	"iter"
)

func ScannerToIter(s bufio.Scanner) iter.Seq[string] {
	return func(yield func(string) bool) {
		for s.Scan() {
			if !yield(s.Text()) {
				break
			}
		}
	}
}
