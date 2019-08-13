package chinumeral

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"unicode/utf8"
	"unsafe"
)

var (
	errNilPointer = errors.New("chinumeral: Chinese numeral decode on nil pointer")
)

var (
	cnAr = map[rune]Chinese{
		// 小写数字
		'〇': 0, '一': 1, '二': 2, '三': 3, '四': 4,
		'五': 5, '六': 6, '七': 7, '八': 8, '九': 9,
		'十': 10, '百': 100, '千': 1e3, '万': 1e4, '亿': 1e8,

		// 大写数字
		'零': 0, '壹': 1, '贰': 2, '叁': 3, '肆': 4,
		'伍': 5, '陆': 6, '柒': 7, '捌': 8, '玖': 9,
		'拾': 10, '佰': 100, '仟': 1e3, '萬': 1e4, '億': 1e8,

		// 中文 阿拉伯数字
		'０': 0, '１': 1, '２': 2, '３': 3, '４': 4,
		'５': 5, '６': 6, '７': 7, '８': 8, '９': 9,

		// ascii 阿拉伯数字
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4,
		'5': 5, '6': 6, '7': 7, '8': 8, '9': 9,

		// 特殊的
		'幺': 1,
		'两': 2, '兩': 2,
		'卄': 20, '廿': 20,
		'卅': 30,
		'卌': 40,
	}
)

var (
	CHT = &ChineseFormat{
		Basic: [10]string{
			"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖",
		},
		Carry10: [3]string{
			"拾", "佰", "仟",
		},
		Carry1e4: [2]string{
			"萬", "億",
		},
	}

	CHS = &ChineseFormat{
		Basic: [10]string{
			"零", "一", "二", "三", "四", "五", "六", "七", "八", "九",
		},
		Carry10: [3]string{
			"十", "百", "千",
		},
		Carry1e4: [2]string{
			"万", "亿",
		},
	}
)

type ChineseFormat struct {
	// 零 至 九
	Basic [10]string

	// 十 百 千
	Carry10 [3]string

	// 万 亿
	Carry1e4 [2]string
}

type Chinese uint64

func (c *Chinese) Decode(s []byte) (n int, err error) {
	if c == nil {
		return 0, errNilPointer
	}
	var result Chinese
	var tmp Chinese
	var mln Chinese

	for n != len(s) {
		ch, size := utf8.DecodeRune(s[n:])
		curr, ok := cnAr[ch]
		if !ok {
			break
		}
		switch curr {
		case 1e8:
			mln = (mln + result + tmp) * curr
			result = 0
			tmp = 0
		case 1e4:
			result = (result + tmp) * curr
			tmp = 0
		case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
			tmp = tmp*10 + curr
		default:
			if tmp != 0 {
				curr *= tmp
			}
			result += curr
			tmp = 0
		}
		n += size
	}
	*c = result + tmp + mln
	return n, nil
}

func (c *Chinese) DecodeString(s string) (n int, err error) {
	return c.Decode(*(*[]byte)(unsafe.Pointer(&s)))
}

func (c Chinese) encodeZreo(w io.Writer) (err error) {
	_, err = io.WriteString(w, CHS.Basic[0])
	if err != nil {
		return err
	}
	return nil
}

func (c Chinese) encodeToWriter(w io.Writer, opt *ChineseFormat) (err error) {
	if c == 0 {
		return c.encodeZreo(w)
	}
	for c != 0 {
		switch {
		case c >= 1e8:
			err = (c / 1e8).encodeToWriter(w, opt)
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, opt.Carry1e4[1])
			if err != nil {
				return err
			}
			c %= 1e8
			if c != 0 && c < 1e7 {
				err = c.encodeZreo(w)
				if err != nil {
					return err
				}
			}
		case c >= 1e4:
			err = (c / 1e4).encodeToWriter(w, opt)
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, opt.Carry1e4[0])
			if err != nil {
				return err
			}
			c %= 1e4
			if c != 0 && c < 1e3 {
				err = c.encodeZreo(w)
				if err != nil {
					return err
				}
			}
		case c >= 1e3:
			err = (c / 1e3).encodeToWriter(w, opt)
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, opt.Carry10[2])
			if err != nil {
				return err
			}
			c %= 1e3
			if c != 0 && c < 1e2 {
				err = c.encodeZreo(w)
				if err != nil {
					return err
				}
			}
		case c >= 1e2:
			err = (c / 1e2).encodeToWriter(w, opt)
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, opt.Carry10[1])
			if err != nil {
				err = c.encodeZreo(w)
				if err != nil {
					return err
				}
			}
			c %= 1e2
			if c != 0 && c < 10 {
				err = c.encodeZreo(w)
				if err != nil {
					return err
				}
			}
		case c >= 10:
			err = (c / 10).encodeToWriter(w, opt)
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, opt.Carry10[0])
			if err != nil {
				return err
			}
			c %= 10
		default:
			_, err = io.WriteString(w, opt.Basic[c])
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}

func (c Chinese) Encode(opt *ChineseFormat) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := c.encodeToWriter(buf, opt)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c Chinese) EncodeToString(opt *ChineseFormat) (string, error) {
	b, err := c.Encode(opt)
	if err != nil {
		return "", err
	}
	return *(*string)(unsafe.Pointer(&b)), nil
}

func (c Chinese) String() string {
	ret, err := c.EncodeToString(CHS)
	if err != nil {
		return fmt.Sprintf("Chinese(%d)", uint64(c))
	}
	return ret
}
