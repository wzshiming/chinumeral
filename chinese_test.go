package chinumeral

import (
	"math"
	"testing"
)

func TestChinese(t *testing.T) {
	opts := []*ChineseOption{
		Upper, Lower, Number,
	}
	ranges := [][2]Chinese{
		{0, 1e6 + 1e5},
		{1e7 - 1e5, 1e7 + 1e5},
		{1e8 - 1e5, 1e8 + 1e5},
		{1e9 - 1e5, 1e9 + 1e5},
		{1e10 - 1e5, 1e10 + 1e5},
		{1e11 - 1e5, 1e11 + 1e5},
		{1e12 - 1e5, 1e12 + 1e5},
		{1e13 - 1e5, 1e13 + 1e5},
		{1e14 - 1e5, 1e14 + 1e5},
		{math.MaxUint64 - 1e5, math.MaxUint64},
	}

	for _, opt := range opts {
		for _, ran := range ranges {
			for i := ran[0]; i != ran[1]; i++ {
				tmp, err := i.EncodeToString(opt)
				if err != nil {
					t.Error(err)
				}
				var d Chinese
				_, err = d.DecodeString(tmp)
				if err != nil {
					t.Error(err)
				}
				if i != d {
					t.Fatal(uint64(i), tmp, uint64(d))
				}
			}
		}
	}
}
