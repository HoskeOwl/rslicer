package rslicer

import (
	"errors"
	"fmt"
	"testing"
)

var (
	str  = "abcdefgh"
	ustr = "猫japが食wordsべる"
)

type sRange struct {
	begin, end int
}

func (s sRange) String() string {
	return fmt.Sprintf("%v:%v", s.begin, s.end)
}

type tData struct {
	str       string
	runeRange sRange
	expRange  sRange
}

func (t *tData) ExpectedString() string {
	return string([]rune(t.str)[t.expRange.begin:t.expRange.end])
}

func (t tData) String() string {
	return fmt.Sprintf("str: '%v', runeRange: %v, expected: %v, expectedRange: %v", t.str, t.runeRange, t.ExpectedString(), t.expRange)
}

func TestRangeSliceOk(T *testing.T) {
	var check []tData = []tData{
		// ascii
		tData{str, sRange{0, 1}, sRange{0, 1}}, // "a"
		tData{str, sRange{0, 2}, sRange{0, 2}}, // "ab"
		tData{str, sRange{0, 7}, sRange{0, 7}}, // "abcdefg"
		tData{str, sRange{0, 8}, sRange{0, 8}}, // "abcdefgh"
		tData{str, sRange{1, 8}, sRange{1, 8}}, // "bcdefgh"
		tData{str, sRange{7, 8}, sRange{7, 8}}, // "h"

		tData{str, sRange{0, -1}, sRange{0, 7}},  // "abcdefg"
		tData{str, sRange{1, -1}, sRange{1, 7}},  // "bcdefg"
		tData{str, sRange{-7, -1}, sRange{1, 7}}, // "bcdefg"

		// unicode
		tData{ustr, sRange{0, 1}, sRange{0, 1}},   // "猫"
		tData{ustr, sRange{0, 3}, sRange{0, 3}},   // "猫ja"
		tData{ustr, sRange{1, 5}, sRange{1, 5}},   // "japが食"
		tData{ustr, sRange{0, 13}, sRange{0, 13}}, // "猫japが食wordsべる"
		tData{ustr, sRange{1, 13}, sRange{1, 13}}, // "japが食wordsべる"
		tData{ustr, sRange{1, 12}, sRange{1, 12}}, // "japが食wordsべ"

		tData{ustr, sRange{0, -1}, sRange{0, 12}},   // "猫japが食wordsべ"
		tData{ustr, sRange{1, -1}, sRange{1, 12}},   // "japが食words"
		tData{ustr, sRange{-12, -1}, sRange{1, 12}}, // "japが食words"
	}

	for _, tdata := range check {
		b, e, err := GetRuneRange(tdata.str, tdata.runeRange.begin, tdata.runeRange.end)
		if err != nil {
			T.Errorf("Got error: %v. %v", err, tdata)
		}
		got := tdata.str[b:e]
		expected := tdata.ExpectedString()

		if got != expected {
			T.Errorf("Wrong slice, '%v' != '%v', %v", got, expected, tdata)
		}
	}
}

func TestSliceOk(T *testing.T) {
	var check []tData = []tData{
		// ascii
		tData{str, sRange{0, 1}, sRange{0, 1}}, // "a"
		tData{str, sRange{0, 2}, sRange{0, 2}}, // "ab"
		tData{str, sRange{0, 7}, sRange{0, 7}}, // "abcdefg"
		tData{str, sRange{0, 8}, sRange{0, 8}}, // "abcdefgh"
		tData{str, sRange{1, 8}, sRange{1, 8}}, // "bcdefgh"
		tData{str, sRange{7, 8}, sRange{7, 8}}, // "h"

		tData{str, sRange{0, -1}, sRange{0, 7}},  // "abcdefg"
		tData{str, sRange{1, -1}, sRange{1, 7}},  // "bcdefg"
		tData{str, sRange{-7, -1}, sRange{1, 7}}, // "bcdefg"

		// unicode
		tData{ustr, sRange{0, 1}, sRange{0, 1}},   // "猫"
		tData{ustr, sRange{0, 3}, sRange{0, 3}},   // "猫ja"
		tData{ustr, sRange{1, 5}, sRange{1, 5}},   // "japが食"
		tData{ustr, sRange{0, 13}, sRange{0, 13}}, // "猫japが食wordsべる"
		tData{ustr, sRange{1, 13}, sRange{1, 13}}, // "japが食wordsべる"
		tData{ustr, sRange{1, 12}, sRange{1, 12}}, // "japが食wordsべ"

		tData{ustr, sRange{0, -1}, sRange{0, 12}},   // "猫japが食wordsべ"
		tData{ustr, sRange{1, -1}, sRange{1, 12}},   // "japが食words"
		tData{ustr, sRange{-12, -1}, sRange{1, 12}}, // "japが食words"
	}

	for _, tdata := range check {
		got, err := GetSliceByRunes(tdata.str, tdata.runeRange.begin, tdata.runeRange.end)
		if err != nil {
			T.Errorf("Got error: %v. %v", err, tdata)
		}
		expected := tdata.ExpectedString()

		if got != expected {
			T.Errorf("Wrong slice, '%v' != '%v', %v", got, expected, tdata)
		}
	}
}

type eData struct {
	str       string
	runeRange sRange
}

func (e eData) String() string {
	return fmt.Sprintf("str: '%v', runeRange: %v", e.str, e.runeRange)
}

func TestRangeSliceErr(T *testing.T) {
	var check []eData = []eData{
		// ascii
		eData{str, sRange{7, 1}},
		eData{str, sRange{1, 112}},
		eData{str, sRange{120, 150}},
		eData{str, sRange{120, 3}},

		eData{str, sRange{-5, -5}},
		eData{str, sRange{0, -15}},
		eData{str, sRange{-15, -1}},
		eData{str, sRange{-17, -18}},

		// unicode
		eData{ustr, sRange{7, 1}},
		eData{ustr, sRange{1, 31}},
		eData{ustr, sRange{120, 155}},
		eData{ustr, sRange{120, 3}},

		eData{ustr, sRange{-5, -5}},
		eData{ustr, sRange{0, -15}},
		eData{ustr, sRange{-15, -1}},
		eData{ustr, sRange{-17, -18}},
	}

	for _, edata := range check {
		_, _, err := GetRuneRange(edata.str, edata.runeRange.begin, edata.runeRange.end)
		if err == nil {
			T.Errorf("Didn't get error: %v. %v", err, edata)
		}
		if !errors.Is(ErrOutOfRange, err) {
			T.Errorf("Wrong error. Got: %v", err)
		}
	}
}

func TestSliceErr(T *testing.T) {
	var check []eData = []eData{
		// ascii
		eData{str, sRange{7, 1}},
		eData{str, sRange{1, 112}},
		eData{str, sRange{120, 150}},
		eData{str, sRange{120, 3}},

		eData{str, sRange{-5, -5}},
		eData{str, sRange{0, -15}},
		eData{str, sRange{-15, -1}},
		eData{str, sRange{-17, -18}},

		// unicode
		eData{ustr, sRange{7, 1}},
		eData{ustr, sRange{1, 31}},
		eData{ustr, sRange{120, 155}},
		eData{ustr, sRange{120, 3}},

		eData{ustr, sRange{-5, -5}},
		eData{ustr, sRange{0, -15}},
		eData{ustr, sRange{-15, -1}},
		eData{ustr, sRange{-17, -18}},
	}

	for _, edata := range check {
		_, err := GetSliceByRunes(edata.str, edata.runeRange.begin, edata.runeRange.end)
		if err == nil {
			T.Errorf("Didn't get error: %v. %v", err, edata)
		}
		if !errors.Is(ErrOutOfRange, err) {
			T.Errorf("Wrong error. Got: %v", err)
		}
	}
}
