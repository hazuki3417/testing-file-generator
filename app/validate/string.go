package validate

import (
	"fmt"

	verify "github.com/hazuki3417/testing-file-generator/verify"
)

type strVerify struct {
	verify verify.StrVerify
}

func StrValidate(value string) strVerify {
	return strVerify{verify: verify.StrVerify{Value: value}}
}

func (str *strVerify) Valid() error {
	if str.verify.Empty() {
		// 空文字ならエラーにする
		return fmt.Errorf("空の文字列は指定できません")
	}
	if str.verify.OnlyInvalidCharacters() {
		// 半角・全角・改行コードのみの文字列ならエラーにする
		return fmt.Errorf("半角、全角、改行コードのみの文字列は指定できません")
	}
	return nil
}

func (sv *strVerify) Between(min int, max int) error {
	if !sv.verify.InRange(min, max) {
		return fmt.Errorf("%v字以上、%v字以下となる文字列を指定してください", min, max)
	}
	return nil
}
