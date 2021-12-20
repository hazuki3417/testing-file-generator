package validate_test

import (
	"errors"
	"testing"

	validate "github.com/hazuki3417/testing-file-generator/validate"
	"github.com/stretchr/testify/assert"
)

func Test_IntValidate_Size(t *testing.T) {

	type testVar struct {
		value int
	}

	type argVar struct {
		min int
		max int
	}

	type expVar struct {
		result error
	}

	// テストパターンを作成
	tests := []struct {
		testName string
		testVar  testVar
		argVar   argVar
		expVar   expVar
	}{
		{
			// minの境界値テスト
			testName: "patten1",
			testVar:  testVar{value: 1},
			argVar:   argVar{min: 2, max: 5},
			expVar:   expVar{result: errors.New("2以上、5以下となる値を指定してください")},
		},
		{
			// minの境界値テスト
			testName: "patten2",
			testVar:  testVar{value: 2},
			argVar:   argVar{min: 2, max: 5},
			expVar:   expVar{result: nil},
		},
		{
			// maxの境界値テスト
			testName: "patten3",
			testVar:  testVar{value: 6},
			argVar:   argVar{min: 2, max: 5},
			expVar:   expVar{result: errors.New("2以上、5以下となる値を指定してください")},
		},
		{
			// maxの境界値テスト
			testName: "patten4",
			testVar:  testVar{value: 5},
			argVar:   argVar{min: 2, max: 5},
			expVar:   expVar{result: nil},
		},
	}
	for _, tp := range tests {
		value := validate.IntValidate(tp.testVar.value)

		t.Run(tp.testName, func(t *testing.T) {
			assert.EqualValues(t, value.Range(tp.argVar.min, tp.argVar.max), tp.expVar.result)
		})
	}
}
