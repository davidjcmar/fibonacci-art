package period

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPisanoSmall(t *testing.T) {
	var tests = []struct {
			modulus uint
			want []uint64 // verify by []uint64 value
		}{
			{2, []uint64{0,1,1}},
			{3, []uint64{0,1,1,2,0,2,2,1}},
		}

		for _, tt := range(tests) {
			testname := fmt.Sprintf("FibPisano mod %v", tt.modulus)
			t.Run(testname, func(t *testing.T) {
				got, _ := Pisano(tt.modulus)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("got %v want %v", got, tt.want)
				}
			})
		}
}

func TestPisanoLarge(t *testing.T) {
	var tests = []struct {
		modulus uint
		want int // verify by len([]uint64)
	}{
		{48, 24},
		{50, 300},
		{101, 50},
		{144, 24},
	}

	for _, tt := range(tests) {
		testname := fmt.Sprintf("FibPisano mod %v", tt.modulus)
		t.Run(testname, func(t *testing.T) {
			got, _ := Pisano(tt.modulus)
			if len(got) != tt.want {
				t.Errorf("got %v want %v", len(got), tt.want)
			}
		})
	}
}

