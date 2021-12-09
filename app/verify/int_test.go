package verify_test

import (
	"testing"

	verify "github.com/hazuki3417/testing-file-generator/verify"
	"github.com/stretchr/testify/assert"
)

func Test_IntVerify(t *testing.T) {

	type testVar = verify.IntVerify

	type testArg struct {
		size int
	}

	type expVar struct {
		equal    bool
		notEqual bool
		gt       bool
		ge       bool
		le       bool
		lt       bool
		negative bool
		positive bool
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
			testVar:  testVar{Value: 7},
			testArg:  testArg{size: 6},
			expVar: expVar{
				equal: false, notEqual: true,
				gt: true, ge: true, le: false, lt: false,
				negative: false, positive: true,
			},
		},
		{
			testName: "patten2",
			testVar:  testVar{Value: 7},
			testArg:  testArg{size: 8},
			expVar: expVar{
				equal: false, notEqual: true,
				gt: false, ge: false, le: true, lt: true,
				negative: false, positive: true,
			},
		},
		{
			testName: "patten3",
			testVar:  testVar{Value: 7},
			testArg:  testArg{size: 7},
			expVar: expVar{
				equal: true, notEqual: false,
				gt: false, ge: true, le: true, lt: false,
				negative: false, positive: true,
			},
		},
		{
			testName: "patten4",
			testVar:  testVar{Value: -1},
			testArg:  testArg{size: 0},
			expVar: expVar{
				equal: false, notEqual: true,
				gt: false, ge: false, le: true, lt: true,
				negative: true, positive: false,
			},
		},
	}

	for _, tp := range tests {
		sv := tp.testVar
		t.Run(tp.testName, func(t *testing.T) {
			assert.EqualValues(t, sv.Equal(tp.testArg.size), tp.expVar.equal)
			assert.EqualValues(t, sv.NotEqual(tp.testArg.size), tp.expVar.notEqual)
			assert.EqualValues(t, sv.Gt(tp.testArg.size), tp.expVar.gt)
			assert.EqualValues(t, sv.Ge(tp.testArg.size), tp.expVar.ge)
			assert.EqualValues(t, sv.Le(tp.testArg.size), tp.expVar.le)
			assert.EqualValues(t, sv.Lt(tp.testArg.size), tp.expVar.lt)
			assert.EqualValues(t, sv.Negative(), tp.expVar.negative)
			assert.EqualValues(t, sv.Positive(), tp.expVar.positive)
		})
	}
}

func Test_IntVerifyRange(t *testing.T) {

	type testVar = verify.IntVerify

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
			testVar:  testVar{Value: 7},
			testArg:  testArg{min: 2, max: 7},
			expVar:   expVar{inRange: true, outRange: false},
		},
		{
			testName: "patten2",
			testVar:  testVar{Value: 7},
			testArg:  testArg{min: 8, max: 10},
			expVar:   expVar{inRange: false, outRange: true},
		},
		{
			testName: "patten3",
			testVar:  testVar{Value: 7},
			testArg:  testArg{min: 10, max: 13},
			expVar:   expVar{inRange: false, outRange: true},
		},
		{
			testName: "patten4",
			testVar:  testVar{Value: 0},
			testArg:  testArg{min: 0, max: 6},
			expVar:   expVar{inRange: true, outRange: false},
		},
	}

	for _, tp := range tests {
		lv := tp.testVar
		t.Run(tp.testName, func(t *testing.T) {
			if ver := lv.InRange(tp.testArg.min, tp.testArg.max); ver != tp.expVar.inRange {
				t.Errorf("InRange() = %v, expected %v", ver, tp.expVar.inRange)
			}
			if ver := lv.OutRange(tp.testArg.min, tp.testArg.max); ver != tp.expVar.outRange {
				t.Errorf("OutRange() = %v, expected %v", ver, tp.expVar.outRange)
			}
		})
	}
}
