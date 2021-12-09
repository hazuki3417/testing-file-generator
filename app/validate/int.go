package validate

import (
	"fmt"

	verify "github.com/hazuki3417/testing-file-generator/verify"
)

type intVerify struct {
	verify verify.IntVerify
}

func IntValidate(value int) intVerify {
	return intVerify{verify: verify.IntVerify{Value: value}}
}

func (integer *intVerify) Range(min int, max int) error {
	if !integer.verify.InRange(min, max) {
		return fmt.Errorf("%v以上、%v以下となる値を指定してください", min, max)
	}
	return nil
}
