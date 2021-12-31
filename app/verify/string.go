package verify

import (
	"regexp"
	"unicode/utf8"
)

// string型の値を検証する構造体です
type StrVerify struct {
	Value string
}

func (sv *StrVerify) Empty() bool {
	return sv.Value == ""
}

func (sv *StrVerify) NotEmpty() bool {
	return sv.Value != ""
}

func (sv *StrVerify) OnlyInvalidCharacters() bool {
	// NOTE: \sと全角文字のみのパターンを探す
	pattern := regexp.MustCompile(`^[\s　]+$`)
	return pattern.MatchString(sv.Value)
}

func (sv *StrVerify) Count() int {
	return utf8.RuneCountInString(sv.Value)
}

func (sv *StrVerify) Equal(size int) bool {
	return size == sv.Count()
}

func (sv *StrVerify) NotEqual(size int) bool {
	return size != sv.Count()
}

func (sv *StrVerify) Gt(size int) bool {
	return size < sv.Count()
}

func (sv *StrVerify) Ge(size int) bool {
	return size <= sv.Count()
}

func (sv *StrVerify) Le(size int) bool {
	return sv.Count() <= size
}

func (sv *StrVerify) Lt(size int) bool {
	return sv.Count() < size
}

func (sv *StrVerify) InRange(min int, max int) bool {
	return min <= sv.Count() && sv.Count() <= max
}

func (sv *StrVerify) OutRange(min int, max int) bool {
	return sv.Count() < min || max < sv.Count()
}
