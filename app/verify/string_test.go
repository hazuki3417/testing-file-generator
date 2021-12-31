package verify_test

import (
	"testing"

	verify "github.com/hazuki3417/testing-file-generator/verify"
	"github.com/stretchr/testify/assert"
)

func Test_StrVerify(t *testing.T) {

	type testVar = verify.StrVerify

	type expVar struct {
		empty                 bool
		notEmpty              bool
		onlyInvalidCharacters bool
	}

	// テストパターンを作成
	tests := []struct {
		testName string
		testVar  testVar
		expVar   expVar
	}{
		{
			testName: "patten01",
			testVar:  testVar{Value: ""},
			expVar:   expVar{empty: true, notEmpty: false, onlyInvalidCharacters: false},
		},
		{
			testName: "patten02",
			testVar:  testVar{Value: "testing"},
			expVar:   expVar{empty: false, notEmpty: true, onlyInvalidCharacters: false},
		},
		{
			// 半角文字が混在するテスト
			testName: "patten03",
			testVar:  testVar{Value: "testing "},
			expVar:   expVar{empty: false, notEmpty: true, onlyInvalidCharacters: false},
		},
		{
			// 半角文字が混在するテスト
			testName: "patten04",
			testVar:  testVar{Value: " testing"},
			expVar:   expVar{empty: false, notEmpty: true, onlyInvalidCharacters: false},
		},
		{
			// 半角文字が混在するテスト
			testName: "patten05",
			testVar:  testVar{Value: "test ing"},
			expVar:   expVar{empty: false, notEmpty: true, onlyInvalidCharacters: false},
		},
		{
			// 全角文字が混在するテスト
			testName: "patten06",
			testVar:  testVar{Value: "testing　"},
			expVar:   expVar{empty: false, notEmpty: true, onlyInvalidCharacters: false},
		},
		{
			// 全角文字が混在するテスト
			testName: "patten07",
			testVar:  testVar{Value: "　testing"},
			expVar:   expVar{empty: false, notEmpty: true, onlyInvalidCharacters: false},
		},
		{
			// 全角文字が混在するテスト
			testName: "patten08",
			testVar:  testVar{Value: "test　ing"},
			expVar:   expVar{empty: false, notEmpty: true, onlyInvalidCharacters: false},
		},
		{
			// 半角文字が1つのテスト
			testName: "patten09",
			testVar:  testVar{Value: " "},
			expVar:   expVar{empty: false, notEmpty: true, onlyInvalidCharacters: true},
		},
		{
			// 半角文字が複数のテスト
			testName: "patten10",
			testVar:  testVar{Value: "  "},
			expVar:   expVar{empty: false, notEmpty: true, onlyInvalidCharacters: true},
		},
		{
			// 全角文字が1つのテスト
			testName: "patten11",
			testVar:  testVar{Value: "　"},
			expVar:   expVar{empty: false, notEmpty: true, onlyInvalidCharacters: true},
		},
		{
			// 全角文字が複数のテスト
			testName: "patten12",
			testVar:  testVar{Value: "　　"},
			expVar:   expVar{empty: false, notEmpty: true, onlyInvalidCharacters: true},
		},
		{
			// 全角・半角文字が複数存在するテスト
			testName: "patten13",
			testVar:  testVar{Value: " 　 　"},
			expVar:   expVar{empty: false, notEmpty: true, onlyInvalidCharacters: true},
		},
	}

	for _, tp := range tests {
		lv := tp.testVar
		t.Run(tp.testName, func(t *testing.T) {
			assert.EqualValues(t, lv.Empty(), tp.expVar.empty)
			assert.EqualValues(t, lv.NotEmpty(), tp.expVar.notEmpty)
			assert.EqualValues(t, lv.OnlyInvalidCharacters(), tp.expVar.onlyInvalidCharacters)
		})
	}
}
func Test_StrVerify_Count(t *testing.T) {

	type testVar = verify.StrVerify

	type expVar struct {
		count int
	}

	// テストパターンを作成
	tests := []struct {
		testName string
		testVar  testVar
		expVar   expVar
	}{
		{
			testName: "patten01",
			testVar:  testVar{Value: ""},
			expVar:   expVar{count: 0},
		},
		{
			// 半角文字のカウントテスト
			testName: "patten02",
			testVar:  testVar{Value: "aaa"},
			expVar:   expVar{count: 3},
		},
		{
			// 全角文字のカウントテスト
			testName: "patten03",
			testVar:  testVar{Value: "あああああ"},
			expVar:   expVar{count: 5},
		},
	}

	for _, tp := range tests {
		lv := tp.testVar
		t.Run(tp.testName, func(t *testing.T) {
			assert.EqualValues(t, lv.Count(), tp.expVar.count)
		})
	}
}

func Test_StrVerify_Length(t *testing.T) {

	type testVar = verify.StrVerify

	type testArg = struct {
		size int
	}

	type expVar struct {
		equal bool
		gt    bool
		ge    bool
		le    bool
		lt    bool
	}

	// テストパターンを作成
	tests := []struct {
		testName string
		testVar  testVar
		testArg  testArg
		expVar   expVar
	}{
		{
			testName: "patten1",
			testVar:  testVar{Value: "testing"},
			testArg:  testArg{size: 6},
			expVar:   expVar{equal: false, gt: true, ge: true, le: false, lt: false},
		},
		{
			testName: "patten2",
			testVar:  testVar{Value: "testing"},
			testArg:  testArg{size: 8},
			expVar:   expVar{equal: false, gt: false, ge: false, le: true, lt: true},
		},
		{
			testName: "patten3",
			testVar:  testVar{Value: "testing"},
			testArg:  testArg{size: 7},
			expVar:   expVar{equal: true, gt: false, ge: true, le: true, lt: false},
		},
		{
			testName: "patten4",
			testVar:  testVar{Value: ""},
			testArg:  testArg{size: 0},
			expVar:   expVar{equal: true, gt: false, ge: true, le: true, lt: false},
		},
	}

	for _, tp := range tests {
		lv := tp.testVar
		t.Run(tp.testName, func(t *testing.T) {
			assert.EqualValues(t, lv.Equal(tp.testArg.size), tp.expVar.equal)
			assert.EqualValues(t, lv.Gt(tp.testArg.size), tp.expVar.gt)
			assert.EqualValues(t, lv.Ge(tp.testArg.size), tp.expVar.ge)
			assert.EqualValues(t, lv.Le(tp.testArg.size), tp.expVar.le)
			assert.EqualValues(t, lv.Lt(tp.testArg.size), tp.expVar.lt)
		})
	}
}

func Test_StrVerify_LengthRange(t *testing.T) {

	type testVar = verify.StrVerify

	type testArg = struct {
		min int
		max int
	}

	type expVar struct {
		inRange  bool
		outRange bool
	}

	// テストパターンを作成
	tests := []struct {
		testName string
		testVar  testVar
		testArg  testArg
		expVar   expVar
	}{
		{
			testName: "patten1",
			testVar:  testVar{Value: "testing"},
			testArg:  testArg{min: 2, max: 7},
			expVar:   expVar{inRange: true, outRange: false},
		},
		{
			testName: "patten2",
			testVar:  testVar{Value: "testing"},
			testArg:  testArg{min: 8, max: 10},
			expVar:   expVar{inRange: false, outRange: true},
		},
		{
			testName: "patten3",
			testVar:  testVar{Value: "testing"},
			testArg:  testArg{min: 10, max: 13},
			expVar:   expVar{inRange: false, outRange: true},
		},
		{
			testName: "patten4",
			testVar:  testVar{Value: ""},
			testArg:  testArg{min: 0, max: 6},
			expVar:   expVar{inRange: true, outRange: false},
		},
	}

	for _, tp := range tests {
		lv := tp.testVar
		t.Run(tp.testName, func(t *testing.T) {
			assert.EqualValues(t, lv.InRange(tp.testArg.min, tp.testArg.max), tp.expVar.inRange)
			assert.EqualValues(t, lv.OutRange(tp.testArg.min, tp.testArg.max), tp.expVar.outRange)
		})
	}
}
