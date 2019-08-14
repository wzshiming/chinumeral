package chinumeral

import (
	"math"
	"testing"
)

func TestChinese(t *testing.T) {
	startd := [][2]Chinese{
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

	for _, s := range startd {
		for i := s[0]; i != s[1]; i++ {
			tmp, err := i.EncodeToString(Lower)
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
