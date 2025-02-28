package gojq_test

import (
	"fmt"
	"math"
	"math/big"
	"testing"

	"github.com/itchyny/gojq"
)

func TestMarshal(t *testing.T) {
	testCases := []struct {
		value    interface{}
		expected string
	}{
		{
			value:    nil,
			expected: "null",
		},
		{
			value:    []interface{}{false, true},
			expected: "[false,true]",
		},
		{
			value: []interface{}{
				42, 3.14, 1e-6, 1e-7, -1e-9, 1e-10, math.NaN(), math.Inf(1), math.Inf(-1),
				new(big.Int).SetBytes([]byte("\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff")),
			},
			expected: "[42,3.14,0.000001,1e-7,-1e-9,1e-10,null,1.7976931348623157e+308,-1.7976931348623157e+308,340282366920938463463374607431768211455]",
		},
		{
			value:    []interface{}{"", "abcde", "foo\x00\x1f\r\n\t\f\b<=>!\"#$%'& \\\x7fbar"},
			expected: `["","abcde","foo\u0000\u001f\r\n\t\f\b<=>!\"#$%'& \\\u007fbar"]`,
		},
		{
			value:    []interface{}{1, []interface{}{2, []interface{}{3, []interface{}{map[string]interface{}{}}}}},
			expected: `[1,[2,[3,[{}]]]]`,
		},
		{
			value:    map[string]interface{}{"x": []interface{}{100}, "y": map[string]interface{}{"z": 42}},
			expected: `{"x":[100],"y":{"z":42}}`,
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.value), func(t *testing.T) {
			got, err := gojq.Marshal(tc.value)
			if err != nil {
				t.Fatal(err)
			}
			if string(got) != tc.expected {
				t.Errorf("expected: %s, got: %s", tc.expected, string(got))
			}
		})
	}
}
