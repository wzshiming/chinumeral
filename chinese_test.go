package chinumeral

import (
	"math"
	"testing"
)

func TestChinese(t *testing.T) {
	for i := Chinese(0); i != 1e6; i++ {
		tmp, err := i.EncodeToString(CHS)
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

	for i := Chinese(0); i != 1e6; i++ {
		j := i + math.MaxUint64
		tmp, err := j.EncodeToString(CHS)
		if err != nil {
			t.Error(err)
		}
		var d Chinese
		_, err = d.DecodeString(tmp)
		if err != nil {
			t.Error(err)
		}
		if j != d {
			t.Fatal(uint64(j), tmp, uint64(d))
		}
	}

}
