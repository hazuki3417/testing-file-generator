package validate_test

import (
	"errors"
	"testing"

	validate "github.com/hazuki3417/testing-file-generator/validate"
	"github.com/stretchr/testify/assert"
)

func Test_StrValidate_Valid(t *testing.T) {

	type testVar struct {
		value string
	}

	type expVar struct {
		result error
	}

	// テストパターンを作成
	tests := []struct {
		testName string
		testVar  testVar
		expVar   expVar
	}{
		{
			testName: "patten",
			testVar:  testVar{value: ""},
			expVar:   expVar{result: errors.New("空の文字列は指定できません")},
		},
		{
			testName: "patten",
			testVar:  testVar{value: " "}, // 半角スペース1文字
			expVar:   expVar{result: errors.New("半角、全角、改行コードのみの文字列は指定できません")},
		},
		{
			testName: "patten",
			testVar:  testVar{value: "　"}, // 全角スペース1文字
			expVar:   expVar{result: errors.New("半角、全角、改行コードのみの文字列は指定できません")},
		},
		{
			testName: "patten",
			testVar:  testVar{value: "  "}, // 半角スペース2文字
			expVar:   expVar{result: errors.New("半角、全角、改行コードのみの文字列は指定できません")},
		},
		{
			testName: "patten",
			testVar:  testVar{value: "　　"}, // 全角スペース1文字
			expVar:   expVar{result: errors.New("半角、全角、改行コードのみの文字列は指定できません")},
		},
		{
			testName: "patten",
			testVar:  testVar{value: " 　  "}, // 半角・全角混在
			expVar:   expVar{result: errors.New("半角、全角、改行コードのみの文字列は指定できません")},
		},
		{
			testName: "patten",
			testVar:  testVar{value: "a"}, // 1文字のみ指定
			expVar:   expVar{result: nil},
		},
	}
	for _, tp := range tests {
		value := validate.StrValidate(tp.testVar.value)

		t.Run(tp.testName, func(t *testing.T) {
			assert.EqualValues(t, value.Valid(), tp.expVar.result)
		})
	}
}

func Test_StrValidate_Size(t *testing.T) {

	type testVar struct {
		value string
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
			testName: "patten1",
			testVar:  testVar{value: "a"}, // minの境界値テスト
			argVar:   argVar{min: 2, max: 5},
			expVar:   expVar{result: errors.New("2字以上、5字以下となる文字列を指定してください")},
		},
		{
			testName: "patten2",
			testVar:  testVar{value: "aa"}, // minの境界値テスト
			argVar:   argVar{min: 2, max: 5},
			expVar:   expVar{result: nil},
		},
		{
			testName: "patten3",
			testVar:  testVar{value: "aaaaaa"}, // maxの境界値テスト
			argVar:   argVar{min: 2, max: 5},
			expVar:   expVar{result: errors.New("2字以上、5字以下となる文字列を指定してください")},
		},
		{
			testName: "patten4",
			testVar:  testVar{value: "aaaaa"}, // maxの境界値テスト
			argVar:   argVar{min: 2, max: 5},
			expVar:   expVar{result: nil},
		},
	}
	for _, tp := range tests {
		value := validate.StrValidate(tp.testVar.value)

		t.Run(tp.testName, func(t *testing.T) {
			assert.EqualValues(t, value.Between(tp.argVar.min, tp.argVar.max), tp.expVar.result)
		})
	}
}
