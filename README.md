# chinumeral
Chinese numerals in Go


[![Go Report Card](https://goreportcard.com/badge/github.com/wzshiming/chinumeral)](https://goreportcard.com/report/github.com/wzshiming/chinumeral)
[![GoDoc](https://godoc.org/github.com/wzshiming/chinumeral?status.svg)](https://godoc.org/github.com/wzshiming/chinumeral)
[![GitHub license](https://img.shields.io/github/license/wzshiming/chinumeral.svg)](https://github.com/wzshiming/chinumeral/blob/master/LICENSE)
[![gocover.io](https://gocover.io/_badge/github.com/wzshiming/chinumeral)](https://gocover.io/github.com/wzshiming/chinumeral)

## Install

``` shell
go get -u -v github.com/wzshiming/chinumeral
```

## Example

``` golang
Chinese(1000).EncodeToString(Lower)  // 一千
Chinese(1000).EncodeToString(Number) // 一〇〇〇
Chinese(1000).EncodeToString(Upper)  // 壹仟
```

``` golang
func TestChinese(t *testing.T) {
	for i := Chinese(0); i != 1e6; i++ {
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
```

## License

Pouch is licensed under the MIT License. See [LICENSE](https://github.com/wzshiming/chinumeral/blob/master/LICENSE) for the full license text.
