package rslicer

import (
	"unicode/utf8"
)

func getRange(source string, begin, end int) (int, int, error) {
	cur, rbegin, rend := -1, -1, -1
	rcnt := utf8.RuneCountInString(source)
	if (begin < 0 && begin < -rcnt) || (begin > 0 && begin > rcnt) || (end > 0 && end > rcnt) || (end < 0 && end < -rcnt) {
		return 0, 0, ErrOutOfRange
	}
	if begin < 0 {
		begin = rcnt + begin
	}
	if end < 0 {
		end = rcnt + end
	}

	for idx := range source {
		cur += 1
		if cur == begin {
			rbegin = idx
		}
		if cur == end {
			rend = idx
			break
		}
	}
	if rend == -1 {
		rend = len(source)
	}
	return rbegin, rend, nil
}

func GetRuneRange(source string, begin, end int) (rbegin int, rend int, err error) {
	if begin == end || (begin >= 0 && end >= 0 && begin > end) {
		err = ErrOutOfRange
		return
	}
	// If both negative the same sitution
	// begin can't be bigger than end, like "abcdefg"[-1:-3]
	if begin == end || (begin < 0 && end < 0 && begin > end) {
		err = ErrOutOfRange
		return
	}

	rbegin, rend, err = getRange(source, begin, end)
	if rbegin == -1 || rend == -1 {
		rbegin, rend, err = 0, 0, ErrOutOfRange
		return
	}

	return
}
